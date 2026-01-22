package ca

import (
    "bytes"
    "crypto/ecdsa"
    "crypto/elliptic"
    cr "crypto/rand"
    "crypto/tls"
    "crypto/x509"
    "crypto/x509/pkix"
    "encoding/pem"
    "errors"
    "github.com/ethereum/go-ethereum/crypto/ecies"
    _ "log"
    "math/big"
    "math/rand"
    "os"
    "time"
)

type CERT struct {
    CERT       []byte `json:"cert_cert"`
    CERTKEY    *ecdsa.PrivateKey `json:"cert_certkey"`
    CERTPEM    *bytes.Buffer `json:"cert_certpem"`
    CERTKEYPEM *bytes.Buffer `json:"cert_certkeypem"`
    CSR        *x509.Certificate `json:"cert_csr"`
}

func PathExists(path string) (bool,error) {
    _,err := os.Stat(path)
    if err == nil {
        return true,nil
    }
    if os.IsNotExist(err) {
        return false,nil
    }
    return false,err
}

func Req(ca *x509.Certificate, expire int) (string, error) {
    sub := &pkix.Name{
        CommonName:    "china.com",
        Organization:  []string{"test, INC."},
        Country:       []string{"CN"},
        Province:      []string{""},
        Locality:      []string{"test"},
        StreetAddress: []string{"testtest"},
        PostalCode:    []string{"88888"},
    }
    var (
        cert = &CERT{}
        err  error
    )
    cert.CERTKEY, err = ecdsa.GenerateKey(elliptic.P256(),cr.Reader)
    if err != nil {
        return "", err
    }
    if expire < 1 {
        expire = 1
    }
    cert.CSR = &x509.Certificate{
        SerialNumber: big.NewInt(rand.Int63n(2000)),
        Subject:      *sub,
        NotBefore:    time.Now(),
        NotAfter:     time.Now().AddDate(expire, 0, 0),
        SubjectKeyId: []byte{1, 2, 3, 4, 6},
        ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
        KeyUsage:     x509.KeyUsageDigitalSignature,
    }
    cert.CERT, err = x509.CreateCertificate(cr.Reader, cert.CSR, ca, &cert.CERTKEY.PublicKey, cert.CERTKEY)
    if err != nil {
        return "", err
    }
    cert.CERTKEYPEM = new(bytes.Buffer)
    derText,err := x509.MarshalECPrivateKey(cert.CERTKEY)
    if err != nil {
        return "", err
    }
    pem.Encode(cert.CERTKEYPEM, &pem.Block{
        Type:  "ECDSA PRIVATE KEY",
        Bytes: derText,
    })
    return string(cert.CERTKEYPEM.Bytes()), err
}
func ECCEncrypt(pt []byte, puk ecies.PublicKey) ([]byte, error) {
    ct, err := ecies.Encrypt(cr.Reader, &puk, pt, nil, nil)
    return ct, err
}

func Read(buf string) (*ecdsa.PrivateKey, error) {
    // pem解码
    block, _ := pem.Decode([]byte(buf))
    // x509解析
    privateKey,err := x509.ParseECPrivateKey(block.Bytes)
    if err != nil {
        return nil,err
    }
    return privateKey, nil
}

func CreateCA(expire int) (*x509.Certificate, string, error) {
    var (
        ca = new(CERT)
        err error
    )
    //判断是否已经初始化
    f, err := PathExists("./cert/ca/ca.crt")
    if err != nil {
        return nil, "", err
    }
    if f {
        csr, err := LoadPairs("./cert/ca/ca.crt","./cert/ca/ca.key")
        if err != nil {
            return nil, "", err
        }
        return csr, "CA初始化完成！", nil
    }
    sub := &pkix.Name{
        CommonName:    "china.com",
        Organization:  []string{"test, INC."},
        Country:       []string{"CN"},
        Province:      []string{""},
        Locality:      []string{"test"},
        StreetAddress: []string{"testtest"},
        PostalCode:    []string{"88888"},
    }

    if expire < 1 {
        expire = 1
    }
    // 为ca生成私钥
    ca.CERTKEY, err = ecdsa.GenerateKey(elliptic.P256(),cr.Reader)
    if err != nil {
        return nil, "", err
    }

    // 对证书进行签名
    ca.CSR = &x509.Certificate{
        SerialNumber: big.NewInt(rand.Int63n(2000)),
        Subject:      *sub,
        NotBefore:    time.Now(),                       // 生效时间
        NotAfter:     time.Now().AddDate(expire, 0, 0), // 过期时间
        IsCA:         true,                             // 表示用于CA
        ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
        KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
        BasicConstraintsValid: true,
    }
    // 创建证书
    ca.CERT, err = x509.CreateCertificate(cr.Reader, ca.CSR, ca.CSR, &ca.CERTKEY.PublicKey, ca.CERTKEY)
    if err != nil {
        return nil, "", err
    }
    ca.CERTPEM = new(bytes.Buffer)
    pem.Encode(ca.CERTPEM, &pem.Block{
        Type:  "CERTIFICATE",
        Bytes: ca.CERT,
    })
    ca.CERTKEYPEM = new(bytes.Buffer)
    derText,err := x509.MarshalECPrivateKey(ca.CERTKEY)
    if err != nil {
        return nil, "", err
    }
    pem.Encode(ca.CERTKEYPEM, &pem.Block{
        Type:  "ECDSA PRIVATE KEY",
        Bytes: derText,
    })
    // 进行PEM编码，编码就是直接cat证书里面内容显示的东西
    Write(ca, "./ca/cert/ca/ca")
    return ca.CSR, "CA初始化完成！", nil
}


func LoadPairs(certFile, keyFile string) (cert *x509.Certificate, err error) {
    if len(certFile) == 0 && len(keyFile) == 0 {
        return nil, errors.New("cert or key has not provided")
    }
    // 载入cert 和 key文件
    tlsCert, err := tls.LoadX509KeyPair(certFile, keyFile)
    if err != nil {
        return
    }
    cert, err = x509.ParseCertificate(tlsCert.Certificate[0])
    return
}

func Write(cert *CERT, file string) error {
    keyFileName := file + ".key"
    certFIleName := file + ".crt"
    kf, err := os.Create(keyFileName)
    if err != nil {
        return err
    }
    defer kf.Close()

    if _, err := kf.Write(cert.CERTKEYPEM.Bytes()); err != nil {
        return err
    }

    cf, err := os.Create(certFIleName)
    if err != nil {
        return err
    }
    if _, err := cf.Write(cert.CERTPEM.Bytes()); err != nil {
        return err
    }
    return nil
}