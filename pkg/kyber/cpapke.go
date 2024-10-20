package kyber

type CPAPKE struct {
	// contains filtered or unexported fields
}

// KeyGen generates a public key and a secret key for the CCAKEM scheme.
func (kem *CPAPKE) KeyGen() (publicKey, secretKey []byte) {
	// ...
}

func (kem *CPAPKE) Encrypt(publicKey []byte, message []byte, coins []byte) (ciphertext []byte) {
	// ...
}

func (kem *CPAPKE) Decrypt(ciphertext []byte, secretKey []byte) (message []byte) {
	// ...
}
