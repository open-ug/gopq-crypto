package kyber

type CCAKEM struct {
	// contains filtered or unexported fields
}

// KeyGen generates a public key and a secret key for the CCAKEM scheme.
func (kem *CCAKEM) KeyGen() (publicKey, secretKey []byte) {
	// ...
}

func (kem *CCAKEM) Encapsulate(publicKey []byte, coins []byte) (ciphertext, sharedSecret []byte) {
	// ...
}

func (kem *CCAKEM) Decapsulate(ciphertext []byte, secretKey []byte) (sharedSecret []byte) {
	// ...
}
