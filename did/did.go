package did

import (
	"github.com/assetcloud/chain/common"
	"github.com/assetcloud/chain/common/address"
	"github.com/assetcloud/chain/common/crypto"
	_ "github.com/assetcloud/chain/system/address"
	"github.com/assetcloud/chain/system/crypto/secp256k1"
	"github.com/assetcloud/chain/types"
	"github.com/assetcloud/chain/wallet/bipwallet"
)

const (
	MnemonicEnglist = 0
	MnemonicChinese = 1
	AddressType     = 2 // eth format 0xaaaa
)

//NewMnemonic 生成助记词
func NewMnemonic(language int) (string, error) {
	//创建助记词
	return bipwallet.NewMnemonicString(language, 160)

}

//MnemonicToPrivateKey 从助记词到私钥
func MnemonicToPrivateKey(mnemonic string, intdex int) (string, string, error) {
	//创建钱包
	wallet, err := bipwallet.NewWalletFromMnemonic(bipwallet.TypeAS, uint32(types.SECP256K1), mnemonic)
	if err != nil {
		return "", "", err
	}
	//从钱包生成公私钥
	private, public, err := wallet.NewKeyPair(0)
	if err != nil {
		return "", "", err
	}
	return common.ToHex(private), common.ToHex(public), nil
}

//PrivateKeyToPublicKey 从私钥到公钥
func PrivateKeyToPublicKey(privateHex string) (string, error) {
	private, err := common.FromHex(privateHex)
	if err != nil {
		return "", err
	}
	public, err := bipwallet.PrivkeyToPub(bipwallet.TypeAS, uint32(types.SECP256K1), private)
	if err != nil {
		return "", err
	}
	return common.ToHex(public), nil
}

//TestPublicKeyToAddress 从公钥到地址
func PublicKeyToAddress(pubkey string) (string, error) {
	public, err := common.FromHex(pubkey)
	if err != nil {
		return "", err
	}
	address := address.PubKeyToAddr(AddressType, public)
	return address, nil
}

//TestSignature 签名
func Signature(privateHex, msg string) ([]byte, error) {
	c, err := crypto.Load(secp256k1.Name, -1)
	if err != nil {
		return nil, err
	}

	private, err := common.FromHex(privateHex)
	if err != nil {
		return nil, err
	}
	ecdsaPrivate, err := c.PrivKeyFromBytes(private)
	if err != nil {
		return nil, err
	}

	sig := ecdsaPrivate.Sign([]byte(msg))
	return sig.Bytes(), nil
}

//TestVerify 验签
func Verify(publicHex string, sig []byte, msg string) (bool, error) {
	c, err := crypto.Load(secp256k1.Name, -1)
	if err != nil {
		return false, err
	}

	//publicKey: 0x03f9031501ccfc7aeac0a158177ddde9d2da0efe3ef1e732ec80c99f871b8042a2
	//sig: 0x3045022100ec18bf91596195ef0344fdf876012c726091e77bb0119002f22432a28c359fbb02200364e9fc94a458255c06bc7bd1401378599bbce41013472122cecb23e47269c8
	//msg: hello
	//sigHex := "0x3045022100ec18bf91596195ef0344fdf876012c726091e77bb0119002f22432a28c359fbb02200364e9fc94a458255c06bc7bd1401378599bbce41013472122cecb23e47269c8"
	//publicHex := "0x03f9031501ccfc7aeac0a158177ddde9d2da0efe3ef1e732ec80c99f871b8042a2"
	//msg := "hello"

	//sigData, err := common.FromHex(sigHex)

	public, err := common.FromHex(publicHex)
	if err != nil {
		return false, err
	}
	ecdsaPublic, err := c.PubKeyFromBytes(public)
	if err != nil {
		return false, err
	}
	signature, err := c.SignatureFromBytes(sig)
	if err != nil {
		return false, err
	}

	ok := ecdsaPublic.VerifyBytes([]byte(msg), signature)
	return ok, nil
}
