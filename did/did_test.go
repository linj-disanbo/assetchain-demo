package did

import (
	"testing"

	"github.com/assetcloud/chain/common"
	"github.com/stretchr/testify/assert"
)

//TestNewMnemonic 生成助记词
func TestNewMnemonic(t *testing.T) {
	//创建助记词
	mne, err := NewMnemonic(MnemonicChinese)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(mne)
}

//TestMnemonicToPrivateKey 从助记词到私钥
func TestMnemonicToPrivateKey(t *testing.T) {
	//mnemonic: 熟 特 地 罚 坦 皆 礼 函 召 延 多 难 焰 埔 像
	mne := "熟 特 地 罚 坦 皆 礼 函 召 延 多 难 焰 埔 像"
	index := 1
	//创建钱包
	privkey, pubkey, err := MnemonicToPrivateKey(mne, index)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("私钥：%s\n", privkey)
	t.Logf("公钥：%s\n", pubkey)
}

//TestPrivateKeyToPublicKey 从私钥到公钥
func TestPrivateKeyToPublicKey(t *testing.T) {
	//privateKey: 0x0445b197b935d99feaac690a567fb7e84ded0940bc046248213a211dd912285e
	privateHex := "0x0445b197b935d99feaac690a567fb7e84ded0940bc046248213a211dd912285e"
	pubkey, err := PrivateKeyToPublicKey(privateHex)
	assert.Nil(t, err)
	assert.Equal(t, "0x03f9031501ccfc7aeac0a158177ddde9d2da0efe3ef1e732ec80c99f871b8042a2", pubkey)
}

//TestPublicKeyToAddress 从公钥到地址
func TestPublicKeyToAddress(t *testing.T) {
	//publicKey: 0x03f9031501ccfc7aeac0a158177ddde9d2da0efe3ef1e732ec80c99f871b8042a2
	publicHex := "0x03f9031501ccfc7aeac0a158177ddde9d2da0efe3ef1e732ec80c99f871b8042a2"
	address, err := PublicKeyToAddress(publicHex)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, "0x71c31b435d3de6373a0f5a4c40275d0e81dc5ac9", address)
	t.Logf("地址：%s\n", address)
}

//TestSignatureAndVerify 签名  验签
func TestSignatureAndVerify(t *testing.T) {
	//privateKey: 0x0445b197b935d99feaac690a567fb7e84ded0940bc046248213a211dd912285e
	//msg: hello
	privateHex := "0x0445b197b935d99feaac690a567fb7e84ded0940bc046248213a211dd912285e"
	publicHex := "0x03f9031501ccfc7aeac0a158177ddde9d2da0efe3ef1e732ec80c99f871b8042a2"
	msg := "hello"
	hexMsg := common.ToHex([]byte(msg))
	hexMsg = "0x6aa5eea9b5bd86da99e05853e677be2f5f186bf45f2e9273a257b121b3b81f11"
	msg2, _ := common.FromHex(hexMsg)

	assert.Equal(t, "0x68656c6c6f", hexMsg)
	msg = string(msg2)
	sig, err := Signature(privateHex, msg)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("签名结果：%s\n", common.ToHex(sig))
	assert.Equal(t, "0x3045022100cee8338691b6b2be2287a7668c3afa6211bba1f3b7344165fccb598087b5e1e002203be7e90b7fc044fc34fd7aebacb4cebf985a33122304ae6319e759ea46a1f366", common.ToHex(sig))
	ok, err := Verify(publicHex, sig, msg)
	if err != nil {
		t.Error(err)
		return
	}
	assert.Equal(t, true, ok)
	t.Logf("验签结果：%v\n", ok)
}
