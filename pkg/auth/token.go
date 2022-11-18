package auth

import (
	"crypto/tls"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/internal/status"
)


var (
  errMissingMetadate = status.Errorf(codes.InvalidArgument, "missing metadata")
  errInvalidToken = status.Errorf(codes.Unauthenticated, "invalid token")
)

type Auth struct {
  crtFile string `json:"crtFile"`
  keyFile string `json:"keyFile"`
}

func (a *Auth) GetCrtFile() tls.Certificate {
  if cert, err := tls.LoadX509KeyPair(a.crtFile, a.keyFile); err != nil {
    panic(err)
  }

  return cert
}

func (a *Auth) GetTLSServer(cert *tls.Certificate) credentials.TransportCredentials {
  return credentials.NewServerTLSFromCert(cert)
}
