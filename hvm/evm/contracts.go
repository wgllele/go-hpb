// Copyright 2018 The go-hpb Authors
// Modified based on go-ethereum, which Copyright (C) 2014 The go-ethereum Authors.
//
// The go-hpb is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-hpb is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-hpb. If not, see <http://www.gnu.org/licenses/>.

package evm

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math/big"

	"github.com/hpb-project/go-hpb/common"
	"github.com/hpb-project/go-hpb/common/crypto"
	"github.com/hpb-project/go-hpb/common/crypto/bn256"
	"github.com/hpb-project/go-hpb/common/log"
	"github.com/hpb-project/go-hpb/common/math"
	"github.com/hpb-project/go-hpb/config"
	"golang.org/x/crypto/ripemd160"
)

// PrecompiledContract is the basic interface for native Go contracts. The implementation
// requires a deterministic gas count based on the input size of the Run method of the
// contract.
type PrecompiledContract interface {
	RequiredGas(input []byte) uint64  // RequiredPrice calculates the contract gas use
	Run(input []byte) ([]byte, error) // Run runs the precompiled contract
}

// PrecompiledContractsByzantium contains the default set of pre-compiled Hpb
// contracts used in the Byzantium release.
var PrecompiledContractsByzantium = map[common.Address]PrecompiledContract{
	common.BytesToAddress([]byte{1}): &ecrecover{},
	common.BytesToAddress([]byte{2}): &sha256hash{},
	common.BytesToAddress([]byte{3}): &ripemd160hash{},
	common.BytesToAddress([]byte{4}): &dataCopy{},
	common.BytesToAddress([]byte{5}): &bigModExp{},
	common.BytesToAddress([]byte{6}): &bn256Add{},
	common.BytesToAddress([]byte{7}): &bn256ScalarMul{},
	common.BytesToAddress([]byte{8}): &bn256Pairing{},
	common.BytesToAddress([]byte{9}): &zscverify{},
}

// RunPrecompiledContract runs and evaluates the output of a precompiled contract.
func RunPrecompiledContract(p PrecompiledContract, input []byte, contract *Contract) (ret []byte, err error) {
	gas := p.RequiredGas(input)
	if contract.UseGas(gas) {
		return p.Run(input)
	}
	return nil, ErrOutOfGas
}

type G1Point struct {
	x [32]byte
	y [32]byte
}

type G1Pointb struct {
	x []byte
	y []byte
}

func gs(i uint64) G1Point {
	var m [32]byte
	var n [32]byte
	var x []byte
	var y []byte
	if i == 0 {
		a := "0d1fff31f8dfb29333568b00628a0f92a752e8dee420dfede1be731810a807b9"
		b := "06c3001c74387dae9deddc75b76959ef5f98f1be48b0d9fc8ff6d7d76106b41b"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 1 {
		a := "06e1b58cb1420e3d12020c5be2c4e48955efc64310ab10002164d0e2a767018e"
		b := "229facdebea78bd67f5b332bcdab7d692d0c4b18d77e92a8b3ffaee450c797c7"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 2 {
		a := "22f32c65b43f3e770b793ea6e31c85d1aea2c41ea3204fc08a036004e5adef3a"
		b := "1d63e3737f864f05f62e2be0a6b7528b76cdabcda9703edc304c015480fb5543"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 3 {
		a := "01df5e3e2818cfce850bd5d5f57872abc34b1315748e0280c4f0d3d6a40f94a9"
		b := "0d622581880ddba6a3911aa0df64f4fd816800c6dee483f07aa542a6e61534d5"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 4 {
		a := "18d7f2117b1144f5035218384d817c6d1b4359497489a52bcf9d16c44624c1d0"
		b := "115f00d2f27917b5a3e8e6754451a4e990931516cf47e742949b8cbdda0e2c20"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 5 {
		a := "093a9e9ba588d1b8eae48cf96b97def1fb8dccd519678520314e96d289ad1d11"
		b := "0f94a152edd0254ece896bc7e56708ba623c1ed3a27e4fd4c449f8e98fee1b5e"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 6 {
		a := "0a7e8bc3cecaff1d9ec3e7d9c1fab7b5397bd6b6739c99bfe4bcb21d08d25934"
		b := "18d0114fa64774f712044e9a05b818fea4734db2b91fc7f049e120ce01c096be"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 7 {
		a := "2095c16aea6e127aa3394d0124b545a45323708ae1c227575270d99b9900673a"
		b := "24c5a6afc36ef443197217591e084cdd69820401447163b5ab5f015801551a03"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 8 {
		a := "041ee7d5aa6e191ba063876fda64b87728fa3ed39531400118b83372cbb5af75"
		b := "2dc2abc7d618ae4e1522f90d294c23627b6bc4f60093e8f07a7cd3869dac9836"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 9 {
		a := "16dc75831b780dc5806dd5b8973f57f2f4ce8ad2a6bb152fbd9ccb58534115b4"
		b := "17b434c3b65a2f754c99f7bacf2f20bdcd7517a38e5eb301d2d88fe7735ebc9c"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 10 {
		a := "18f1393a76e0af102ffeb380787ed950dc35b04b0cc6de1a6d806d4007b30dba"
		b := "1d640e43bab253bf176b69dffdb3ffc02640c591c392f400596155c8c3f668ef"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 11 {
		a := "2bf3f58b4c957a8ae697aa57eb3f7428527fcb0c7e8d099efae80b97bde600e0"
		b := "14072f8bfdbe285b203cd0a2ebc1aed9ad1de309794226aee63c89397b187abf"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 12 {
		a := "028eb6852c2827302aeb09def685b57bef74ff1a3ff72eda972e32b9ea80c32f"
		b := "1ba2dfb85a585de4b8a189f7b764f87c6f8e06c10d68d4493fc469504888837d"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 13 {
		a := "19003e6b8f14f3583435527eac51a460c705dc6a042a2b7dd56b4f598af50886"
		b := "10e755ac3373f769e7e092f9eca276d911cd31833e82c70b8af09787e2c02d20"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 14 {
		a := "0d493d4d49aa1a4fdf3bc3ba6d969b3e203741b3d570dbc511dd3171baf96f85"
		b := "1d103731795bcc57ddb8514e0e232446bfd9834f6a8ae9ff5235330d2a9e5ffa"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 15 {
		a := "0ce438e766aae8c59b4006ee1749f40370fe5ec9fe29edce6b98e945915db97f"
		b := "02dba20dff83b373d2b47282e08d2c7883254a56701f2dbeea7ccc167ffb49a5"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 16 {
		a := "05092110319650610a94fa0f9d50536404ba526380fc31b99ce95fbc1423a26f"
		b := "18a40146a4e79c2830d6d6e56314c538b0da4a2a72b7533e63f7d0a7e5ab2d22"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 17 {
		a := "25b9ad9c4235b0a2e9f1b2ed20a5ca63814e1fb0eb95540c6f4f163c1a9fc2bd"
		b := "0a726ff7b655ad45468bcfd2d77f8aa0786ff3012d4edb77b5118f863dcdcbc0"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 18 {
		a := "291ff28fa0a9840e230de0f0da725900bd18ce31d2369ffc80abbc4a77c1aff3"
		b := "1ffed5e9dffcd885ac867e2279836a11225548a8c253c47efe24f7d95a4bdd61"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 19 {
		a := "0a01c96340d6bb4c94e028a522f74bef899d8f9d1a6d0b0d832f83275efa68de"
		b := "119c6a17ecb14721ac9eb331abccf2748868855fae43392391c37037d1b150a1"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 20 {
		a := "2c846ad384d3ea063001f34fd60f0b8dc12b3b3ab7a5757f1d394f19850d8309"
		b := "1ff69942134c51e7315ccf1431e66fb5f70c24148c668f4fbe3861fbe535e39c"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 21 {
		a := "0dafb5ae6accb6048e6dbc52f455c262dd2876b565792d68189618a3e630ade0"
		b := "236e97c592c19a2f2244f2938021671045787501e5a4a26de3580628ce37eb3b"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 22 {
		a := "10df3e10a8d613058eae3278e2c80c3366c482354260f501447d15797de7378a"
		b := "10b25f7e075c93203ceba523afc44e0d5cd9e45a60b6dc11d2034180c40a004d"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 23 {
		a := "1437b718d075d54da65adccdd3b6f758a5b76a9e5c5c7a13bf897a92e23fcde2"
		b := "0f0b988d70298608d02c73c410dc8b8bb6b95f0dde0dedcd5ea5692f0c07f3ed"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 24 {
		a := "2705c71a95661231956d10845933f43cd973f4626e3a31dbf6287e01a00beb70"
		b := "27d09bd21d44269e2e7c85e1555fd351698eca14686d5aa969cb08e33db6691b"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 25 {
		a := "1614dabf48099c315f244f8763f4b99ca2cef559781bf55e8e4d912d952edb4a"
		b := "16bf2f8fb1021b47be88ceb6fce08bf3b3a17026509cf9756c1a3fbf3b9d70bd"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 26 {
		a := "21c448cfdcf007959812b2c5977cd4a808fa25408547e660c3fc12ed47501eb3"
		b := "14495c361cf9dc10222549bc258a76a20058f4795c2e65cd27f013c940b7dc7b"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 27 {
		a := "1ac35f37ee0bfcb173d513ea7ac1daf5b46c6f70ce5f82a0396e7afac270ff35"
		b := "2f5f4480260b838ffcba9d34396fc116f75d1d5c24396ed4f7e01fd010ab9970"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 28 {
		a := "0caaa12a18563703797d9be6ef74cbfb9e532cd027a1021f34ad337ce231e074"
		b := "2281c11389906c02bb15e995ffd6db136c3cdb4ec0829b88aec6db8dda05d5af"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 29 {
		a := "1f3d91f1dfbbf01002a7e339ff6754b4ad2290493757475a062a75ec44bc3d50"
		b := "207b99884d9f7ca1e2f04457b90982ec6f8fb0a5b2ffd5b50d9cf4b2d850a920"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 30 {
		a := "1fe58e4e4b1d155fb0a97dc9bae46f401edb2828dc4f96dafb86124cba424455"
		b := "01ad0a57feb7eeda4319a70ea56ded5e9fef71c78ff84413399d51f647d55113"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 31 {
		a := "044e80195798557e870554d7025a8bc6b2ee9a05fa6ae016c3ab3b9e97af5769"
		b := "2c141a12135c4d14352fc60d851cdde147270f76405291b7c5d01da8f5dfed4d"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 32 {
		a := "2883d31d84e605c858cf52260183f09d18bd55dc330f8bf12785e7a2563f8da4"
		b := "0e681e5c997f0bb609af7a95f920f23c4be78ded534832b514510518ede888b2"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 33 {
		a := "2cdf5738c2690b263dfdc2b4235620d781bbff534d3363c4f3cfe5d1c67767c1"
		b := "15f4fb05e5facfd1988d61fd174a14b20e1dbe6ac37946e1527261be8742f5cf"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 34 {
		a := "05542337765c24871e053bb8ec4e1baaca722f58b834426431c6d773788e9c66"
		b := "00e64d379c28d138d394f2cf9f0cc0b5a71e93a055bad23a2c6de74b217f3fac"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 35 {
		a := "2efe9c1359531adb8a104242559a320593803c89a6ff0c6c493d7da5832603ab"
		b := "295898b3b86cf9e09e99d7f80e539078d3b5455bba60a5aa138b2995b75f0409"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 36 {
		a := "2a3740ca39e35d23a5107fdae38209eaebdcd70ae740c873caf8b0b64d92db31"
		b := "05bab66121bccf807b1f776dc487057a5adf5f5791019996a2b7a2dbe1488797"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 37 {
		a := "11ef5ef35b895540be39974ac6ad6697ef4337377f06092b6a668062bf0d8019"
		b := "1a42e3b4b73119a4be1dde36a8eaf553e88717cecb3fdfdc65ed2e728fda0782"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 38 {
		a := "245aac96c5353f38ae92c6c17120e123c223b7eaca134658ebf584a8580ec096"
		b := "25ec55531155156663f8ba825a78f41f158def7b9d082e80259958277369ed08"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 39 {
		a := "0fb13a72db572b1727954bb77d014894e972d7872678200a088febe8bd949986"
		b := "151af2ae374e02dec2b8c5dbde722ae7838d70ab4fd0857597b616a96a1db57c"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 40 {
		a := "155fa64e4c8bf5f5aa53c1f5e44d961f688132c8545323d3bdc6c43a83220f89"
		b := "188507b59213816846bc9c763a93b52fb7ae8e8c8cc7549ce3358728415338a4"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 41 {
		a := "28631525d5192140fd4fb04efbad8dcfddd5b8d0f5dc54442e5530989ef5b7fe"
		b := "0ad3a3d4845b4bc6a92563e72db2bc836168a295c56987c7bb1eea131a3760ac"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 42 {
		a := "043b2963b1c5af8e2e77dfb89db7a0d907a40180929f3fd630a4a37811030b6d"
		b := "0721a4b292b41a3d948237bf076aabeedba377c43a10f78f368042ad155a3c91"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 43 {
		a := "14bfb894e332921cf925f726f7c242a70dbd9366b68b50e14b618a86ecd45bd6"
		b := "09b1c50016fff7018a9483ce00b8ec3b6a0df36db21ae3b8282ca0b4be2e283c"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 44 {
		a := "2758e65c03fdb27e58eb300bde8ada18372aa268b393ad5414e4db097ce9492d"
		b := "041f685536314ddd11441a3d7e01157f7ea7e474aae449dbba70c2edc70cd573"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 45 {
		a := "191365dba9df566e0e6403fb9bcd6847c0964ea516c403fd88543a6a9b3fa1f2"
		b := "0ae815170115c7ce78323cbd9399735847552b379c2651af6fc29184e95eef7f"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 46 {
		a := "027a2a874ba2ab278be899fe96528b6d39f9d090ef4511e68a3e4979bc18a526"
		b := "2272820981fe8a9f0f7c4910dd601cea6dd7045aa4d91843d3cf2afa959fbe68"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 47 {
		a := "13feec071e0834433193b7be17ce48dec58d7610865d9876a08f91ea79c7e28d"
		b := "26325544133c7ec915c317ac358273eb2bf2e6b6119922d7f0ab0727e5eb9e64"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 48 {
		a := "08e6096c8425c13b79e6fa38dffcc92c930d1d0bff9671303dbc0445e73c77bc"
		b := "03e884c8dc85f0d80baf968ae0516c1a7927808f83b4615665c67c59389db606"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 49 {
		a := "1217ff3c630396cd92aa13aa6fee99880afc00f47162625274090278f09cbed3"
		b := "270b44f96accb061e9cad4a3341d72986677ed56157f3ba02520fdf484bb740d"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 50 {
		a := "239128d2e007217328aae4e510c3d9fe1a3ef2b23212dfaf6f2dcb75ef08ed04"
		b := "2d5495372c759fdba858b7f6fa89a948eb4fd277bae9aebf9785c86ea3f9c07d"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 51 {
		a := "305747313ea4d7d17bd14b69527094fa79bdc05c3cc837a668a97eb81cffd3d4"
		b := "0aa43bd7ad9090012e12f78ac3cb416903c2e1aabb61161ca261892465b3555d"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 52 {
		a := "267742bd96caad20a76073d5060085103b7d29c88f0a0d842ef610472a1764ef"
		b := "0086485faeedd1ea8f6595b2edaf5f99044864271a178bd33e6d5b73b6d240a0"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 53 {
		a := "00aed2e1ac448b854a44c7aa43cabb93d92316460c8f5eacb038f4cf554dfa01"
		b := "1b2ec095d370b234214a0c68fdfe8da1e06cbfdc5e889e2337ccb28c49089fcf"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 54 {
		a := "06f37ac505236b2ed8c520ea36b0448229eb2f2536465b14e6e115dc810c6e39"
		b := "174db60e92b421e4d59c81e2c0666f7081067255c8e0d775e085278f34663490"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 55 {
		a := "2af094e58a7961c4a1dba0685d8b01dacbb01f0fc0e7a648085a38aa380a7ab6"
		b := "108ade796501042dab10a83d878cf1deccf74e05edc92460b056d31f3e39fd53"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 56 {
		a := "051ec23f1166a446caa4c8ff443470e98e753697fcceb4fbe5a49bf7a2db7199"
		b := "00f938707bf367e519d0c5efcdb61cc5a606901c0fbd4565abeeb5d020081d96"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 57 {
		a := "1132459cf7287884b102467a71fad0992f1486178f7385ef159277b6e800239d"
		b := "257fedb1e126363af3fb3a80a4ad850d43041d64ef27cc5947730901f3019138"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 58 {
		a := "14a571bbbb8d2a442855cde5fe6ed635d91668eded003d7698f9f744557887ea"
		b := "0f65f76e6fa6f6c7f765f947d905b015c3ad077219fc715c2ec40e37607c1041"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)

	}
	if i == 59 {
		a := "0e303c28b0649b95c624d01327a61fd144d29bfed6d3a1cf83216b45b78180cf"
		b := "229975c2e3aaba1d6203a5d94ea92605edb2af04f41e3783ec4e64755eeb1d1b"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)

	}
	if i == 60 {
		a := "05a62a2f1dfe368e81d9ae5fe150b9a57e0f85572194de27f48fec1c5f3b0dad"
		b := "200eb8097c91fe825adb0e3920e6bdff2e40114bd388298b85a0094a9a5bc654"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)

	}
	if i == 61 {
		a := "06545efc18dfc2f444e147c77ed572decd2b58d0668bbaaf0d31f1297cde6b99"
		b := "29ecbbeb81fe6c14279e9e46637ad286ba71e4c4e5da1416d8501e691f9e5bed"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)

	}
	if i == 62 {
		a := "045ce430f0713c29748e30d024cd703a5672633faebe1fd4d210b5af56a50e70"
		b := "0e3ec93722610f4599ffaac0db0c1b2bb446ff5aea5117710c271d1e64348844"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	if i == 63 {
		a := "243de1ee802dd7a3ca9a991ec228fbbfb4973260f905b5106e5f738183d5cacd"
		b := "133d25bb8dc9f54932b9d6ee98e0432676f5278e9878967fbbd8f5dfc46df4f8"
		x, _ = hex.DecodeString(a)
		y, _ = hex.DecodeString(b)
	}
	copy(m[:], x)
	copy(n[:], y)
	return G1Point{m, n}
}

func (g *G1Point) add(a G1Point) G1Point {
	var input []byte
	var res G1Point
	input = append(input, g.x[:]...)
	input = append(input, g.y[:]...)
	input = append(input, a.x[:]...)
	input = append(input, a.y[:]...)
	p, _ := new(bn256Add).Run(input)
	copy(res.x[:], p[0:32])
	copy(res.y[:], p[32:64])
	return res
}
func (g *G1Point) mul(a []byte) G1Point {
	var input []byte
	var res G1Point
	input = append(input, g.x[:]...)
	input = append(input, g.y[:]...)
	var init [32]byte
	input = append(input, init[:]...)
	input = append(input[:32-len(a)], a...)
	p, _ := new(bn256ScalarMul).Run(input)
	copy(res.x[:], p[0:32])
	copy(res.y[:], p[32:64])
	return res
}

// zscverify implemented as a native contract.
type zscverify struct{}

func (c *zscverify) RequiredGas(input []byte) uint64 {
	return config.ZSCverifyGas
}

const (
	hs_length   = 64
	burn_length = 5
	GROUP_ORDER = 0x30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001
	FIELD_ORDER = 0x30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47
)

func (c *zscverify) Run(input []byte) ([]byte, error) {

	log.Error("zscverify", "string run", string(input))
	log.Error("zscverify", "bytes2hex run", common.Bytes2Hex(input), "len", len(common.Bytes2Hex(input)))

	if len(common.Bytes2Hex(input)) < 6080 {
		return common.LeftPadBytes(input[0:1], 32), nil
	}
	var hs [hs_length]G1Point
	var u G1Point
	var p G1Point
	var ls [hs_length]G1Point
	var rs [hs_length]G1Point
	var a [32]byte
	var b [32]byte
	length := 64
	//salt := int64(binary.BigEndian.Uint64(input[length:]))
	salt := input[length : length+32]
	log.Error("zscverify", "salt", common.Bytes2Hex(salt))
	length = length + 64
	for i := 0; i < 32; i++ {
		//for j := 0; j < 32; j++ {
		//hs[i].x[j] = input[length]
		//length++
		copy(hs[i].x[:], input[length:length+32])
		length += 32
		log.Error("zscverify", "hs i", i, "x", common.Bytes2Hex(hs[i].x[:]))
		// //}
		// for k := 0; k < 32; k++ {
		// 	hs[i].y[k] = input[length]
		// 	length++
		// }
		copy(hs[i].y[:], input[length:length+32])
		log.Error("zscverify", "hs i", i, "y", common.Bytes2Hex(hs[i].y[:]))
		length += 32
	}
	log.Error("zscverify", "length", length)
	// for i := 0; i < 32; i++ {
	// 	u.x[i] = input[length]
	// 	length++
	// }
	// for i := 0; i < 32; i++ {
	// 	u.y[i] = input[length]
	// 	length++
	// }
	copy(u.x[:], input[length:length+32])
	length += 32
	log.Error("zscverify", "u.x", common.Bytes2Hex(u.x[:]))

	copy(u.y[:], input[length:length+32])
	length += 32
	log.Error("zscverify", "u.y", common.Bytes2Hex(u.y[:]))

	// for i := 0; i < 32; i++ {
	// 	p.x[i] = input[length]
	// 	length++
	// }
	// for i := 0; i < 32; i++ {
	// 	p.y[i] = input[length]
	// 	length++
	// }
	copy(p.x[:], input[length:length+32])
	length += 32
	log.Error("zscverify", "p.x", common.Bytes2Hex(p.x[:]))

	copy(p.y[:], input[length:length+32])
	length += 32
	log.Error("zscverify", "p.y", common.Bytes2Hex(p.y[:]))
	length += 32
	for i := 0; i < burn_length; i++ {
		// for j := 0; j < 32; j++ {
		// 	ls[i].x[j] = input[length]
		// 	length++
		// }
		// for k := 0; k < 32; k++ {
		// 	ls[i].y[k] = input[length]
		// 	length++
		// }
		copy(ls[i].x[:], input[length:length+32])
		length += 32
		log.Error("zscverify", "hs i", i, "x", common.Bytes2Hex(ls[i].x[:]))
		copy(ls[i].y[:], input[length:length+32])
		log.Error("zscverify", "hs i", i, "y", common.Bytes2Hex(ls[i].y[:]))
		length += 32
	}
	for i := 0; i < burn_length; i++ {
		// for j := 0; j < 32; j++ {
		// 	rs[i].x[j] = input[length]
		// 	length++
		// }
		// for k := 0; k < 32; k++ {
		// 	rs[i].y[k] = input[length]
		// 	length++
		// }
		copy(rs[i].x[:], input[length:length+32])
		length += 32
		log.Error("zscverify", "rs i", i, "x", common.Bytes2Hex(rs[i].x[:]))
		copy(rs[i].y[:], input[length:length+32])
		log.Error("zscverify", "rs i", i, "y", common.Bytes2Hex(rs[i].y[:]))
		length += 32
	}
	log.Error("zscverify", "length", length)
	// for i := 0; i < 32; i++ {
	// 	a[i] = input[length]
	// 	length++
	// }
	copy(a[:], input[length:length+32])
	length += 32
	log.Error("zscverify", "a", common.Bytes2Hex(a[:]))
	// for i := 0; i < 32; i++ {
	// 	b[i] = input[length]
	// 	length++
	// }
	copy(b[:], input[length:length+32])
	length += 32
	log.Error("zscverify", "b", common.Bytes2Hex(b[:]))
	log_n := burn_length
	n := 2 << (uint(log_n) - 1) //math.Pow(float64(2),float(log_n))
	o := new(big.Int).SetBytes(salt[:])
	var challenges [burn_length]*big.Int
	log.Error("run", "log_n", log_n, "n", n, "o", o)

	//hashdata := crypto.Keccak256([]byte("abcde"))
	//log.Error("HASH", "hash(abcde)", hashdata)
	//log.Error("HASH", "hash(abcde)", common.Bytes2Hex(hashdata))
	g_order := common.Hex2Bytes(string("30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001"))
	//f_order := common.Hex2Bytes(string("30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47"))

	group_order := new(big.Int).SetBytes(g_order)
	//field_order := new(big.Int).SetBytes(f_order)
	for i := 0; i < log_n; i++ {
		var input []byte
		var init [32]byte
		input = append(input, init[:]...)
		input = append(input[:32-len(o.Bytes())], o.Bytes()...)
		input = append(input, ls[i].x[:]...)
		input = append(input, ls[i].y[:]...)
		input = append(input, rs[i].x[:]...)
		input = append(input, rs[i].y[:]...)
		bigh := new(big.Int).SetBytes(crypto.Keccak256(input))
		challenges[i] = new(big.Int).Mod(bigh, group_order)
		o = challenges[i]
		//statement.P = statement.P.add(proof.ls[i].mul(ipAuxiliaries.o.exp(2)).add(proof.rs[i].mul(ipAuxiliaries.o.inv().exp(2))));
		tmp := new(big.Int).Exp(o, new(big.Int).Sub(group_order, new(big.Int).SetInt64(2)), group_order)
		tmp0 := new(big.Int).Exp(o, new(big.Int).SetInt64(2), group_order)
		//statement.P = statement.P.add(proof.ls[i].mul(tmp0.Bytes()).add(proof.rs[i].mul(tmp.exp(2))));
		tmp1 := new(big.Int).Exp(tmp, new(big.Int).SetInt64(2), group_order)
		//statement.P = statement.P.add(proof.ls[i].mul(tmp0.Bytes()).add(proof.rs[i].mul(tmp1.Bytes())));
		lsi := ls[i].mul(tmp0.Bytes())
		rsi := rs[i].mul(tmp1.Bytes())
		p = p.add(lsi.add(rsi))

	}
	log.Error("Step 1")
	var otherExponents [64]*big.Int
	otherExponents[0] = new(big.Int).SetInt64(1)
	for i := 0; i < log_n; i++ {
		tmp := new(big.Int).Mul(otherExponents[0], challenges[i])
		otherExponents[0] = new(big.Int).Mod(tmp, group_order)
	}
	log.Error("Step 2")
	var bitSet [64]bool
	otherExponents[0] = new(big.Int).Exp(otherExponents[0], new(big.Int).Sub(group_order, new(big.Int).SetInt64(2)), group_order)

	for i := uint64(0); i < 32; i++ {
		for j := uint64(0); (1<<j)+i < 32; j++ {
			i1 := i + (1 << j)
			if !bitSet[i1] {
				temp := new(big.Int).Exp(challenges[burn_length-1-j], new(big.Int).SetInt64(2), group_order)
				tmp := new(big.Int).Mul(otherExponents[i], temp)
				otherExponents[i1] = new(big.Int).Mod(tmp, group_order)
				bitSet[i1] = true
			}
		}
	}
	log.Error("Step 3")
	var gTemp G1Point
	var hTemp G1Point

	for i := uint64(0); i < uint64(64); i++ {
		gsi := gs(i)
		gTemp = gTemp.add(gsi.mul(otherExponents[i].Bytes()))
		hTemp = hTemp.add(hs[i].mul(otherExponents[uint64(64-1)-i].Bytes()))
	}
	log.Error("Step 4")
	//var a [32]byte
	//var b [32]byte
	ua := new(big.Int).SetBytes(a[:])
	ub := new(big.Int).SetBytes(b[:])
	uc := new(big.Int).Mul(ua, ub)
	gtp := gTemp.mul(a[:])
	htp := gtp.add(hTemp.mul(b[:]))
	calp := htp.add(u.mul(uc.Bytes()))

	log.Error("Step 5")
	if calp.x == p.x && calp.y == p.y {
		log.Error("zscverify res true")
		return common.LeftPadBytes(input[0:1], 32), nil
	}

	// the first byte of pubkey is bitcoin heritage
	log.Error("zscverify", "res false", string(common.LeftPadBytes(input[0:1], 32)))

	return common.LeftPadBytes(input[0:1], 32), nil
}

// ECRECOVER implemented as a native contract.
type ecrecover struct{}

func (c *ecrecover) RequiredGas(input []byte) uint64 {
	return config.EcrecoverGas
}

func (c *ecrecover) Run(input []byte) ([]byte, error) {
	log.Error("ecrecover", "run", common.Bytes2Hex(input))
	const ecRecoverInputLength = 128

	input = common.RightPadBytes(input, ecRecoverInputLength)
	// "input" is (hash, v, r, s), each 32 bytes
	// but for ecrecover we want (r, s, v)

	r := new(big.Int).SetBytes(input[64:96])
	s := new(big.Int).SetBytes(input[96:128])
	v := input[63] - 27

	// tighter sig s values input homestead only apply to tx sigs
	if !allZero(input[32:63]) || !crypto.ValidateSignatureValues(v, r, s, false) {
		return nil, nil
	}
	// v needs to be at the end for libsecp256k1
	pubKey, err := crypto.Ecrecover(input[:32], append(input[64:128], v))
	// make sure the public key is a valid one
	if err != nil {
		return nil, nil
	}
	log.Error("ecrecover", "res", common.Bytes2Hex(common.LeftPadBytes(crypto.Keccak256(pubKey[1:])[12:], 32)))
	// the first byte of pubkey is bitcoin heritage
	return common.LeftPadBytes(crypto.Keccak256(pubKey[1:])[12:], 32), nil
}

// SHA256 implemented as a native contract.
type sha256hash struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
//
// This method does not require any overflow checking as the input size gas costs
// required for anything significant is so high it's impossible to pay for.
func (c *sha256hash) RequiredGas(input []byte) uint64 {
	return uint64(len(input)+31)/32*config.Sha256PerWordGas + config.Sha256BaseGas
}
func (c *sha256hash) Run(input []byte) ([]byte, error) {
	log.Error("sha256hash", "run", string(input))
	h := sha256.Sum256(input)
	log.Error("sha256hash", "res", string(h[:]))
	log.Error("sha256hash", "res", common.Bytes2Hex(h[:]))

	return h[:], nil
}

// RIPMED160 implemented as a native contract.
type ripemd160hash struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
//
// This method does not require any overflow checking as the input size gas costs
// required for anything significant is so high it's impossible to pay for.
func (c *ripemd160hash) RequiredGas(input []byte) uint64 {
	return uint64(len(input)+31)/32*config.Ripemd160PerWordGas + config.Ripemd160BaseGas
}
func (c *ripemd160hash) Run(input []byte) ([]byte, error) {
	ripemd := ripemd160.New()
	ripemd.Write(input)
	return common.LeftPadBytes(ripemd.Sum(nil), 32), nil
}

// data copy implemented as a native contract.
type dataCopy struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
//
// This method does not require any overflow checking as the input size gas costs
// required for anything significant is so high it's impossible to pay for.
func (c *dataCopy) RequiredGas(input []byte) uint64 {
	return uint64(len(input)+31)/32*config.IdentityPerWordGas + config.IdentityBaseGas
}
func (c *dataCopy) Run(in []byte) ([]byte, error) {
	return in, nil
}

// bigModExp implements a native big integer exponential modular operation.
type bigModExp struct{}

var (
	big1      = big.NewInt(1)
	big4      = big.NewInt(4)
	big8      = big.NewInt(8)
	big16     = big.NewInt(16)
	big32     = big.NewInt(32)
	big64     = big.NewInt(64)
	big96     = big.NewInt(96)
	big480    = big.NewInt(480)
	big1024   = big.NewInt(1024)
	big3072   = big.NewInt(3072)
	big199680 = big.NewInt(199680)
)

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *bigModExp) RequiredGas(input []byte) uint64 {
	var (
		baseLen = new(big.Int).SetBytes(getData(input, 0, 32))
		expLen  = new(big.Int).SetBytes(getData(input, 32, 32))
		modLen  = new(big.Int).SetBytes(getData(input, 64, 32))
	)
	if len(input) > 96 {
		input = input[96:]
	} else {
		input = input[:0]
	}
	// Retrieve the head 32 bytes of exp for the adjusted exponent length
	var expHead *big.Int
	if big.NewInt(int64(len(input))).Cmp(baseLen) <= 0 {
		expHead = new(big.Int)
	} else {
		if expLen.Cmp(big32) > 0 {
			expHead = new(big.Int).SetBytes(getData(input, baseLen.Uint64(), 32))
		} else {
			expHead = new(big.Int).SetBytes(getData(input, baseLen.Uint64(), expLen.Uint64()))
		}
	}
	// Calculate the adjusted exponent length
	var msb int
	if bitlen := expHead.BitLen(); bitlen > 0 {
		msb = bitlen - 1
	}
	adjExpLen := new(big.Int)
	if expLen.Cmp(big32) > 0 {
		adjExpLen.Sub(expLen, big32)
		adjExpLen.Mul(big8, adjExpLen)
	}
	adjExpLen.Add(adjExpLen, big.NewInt(int64(msb)))

	// Calculate the gas cost of the operation
	gas := new(big.Int).Set(math.BigMax(modLen, baseLen))
	switch {
	case gas.Cmp(big64) <= 0:
		gas.Mul(gas, gas)
	case gas.Cmp(big1024) <= 0:
		gas = new(big.Int).Add(
			new(big.Int).Div(new(big.Int).Mul(gas, gas), big4),
			new(big.Int).Sub(new(big.Int).Mul(big96, gas), big3072),
		)
	default:
		gas = new(big.Int).Add(
			new(big.Int).Div(new(big.Int).Mul(gas, gas), big16),
			new(big.Int).Sub(new(big.Int).Mul(big480, gas), big199680),
		)
	}
	gas.Mul(gas, math.BigMax(adjExpLen, big1))
	gas.Div(gas, new(big.Int).SetUint64(config.ModExpQuadCoeffDiv))

	if gas.BitLen() > 64 {
		return math.MaxUint64
	}
	return gas.Uint64()
}

func (c *bigModExp) Run(input []byte) ([]byte, error) {
	var (
		baseLen = new(big.Int).SetBytes(getData(input, 0, 32)).Uint64()
		expLen  = new(big.Int).SetBytes(getData(input, 32, 32)).Uint64()
		modLen  = new(big.Int).SetBytes(getData(input, 64, 32)).Uint64()
	)
	if len(input) > 96 {
		input = input[96:]
	} else {
		input = input[:0]
	}
	// Handle a special case when both the base and mod length is zero
	if baseLen == 0 && modLen == 0 {
		return []byte{}, nil
	}
	// Retrieve the operands and execute the exponentiation
	var (
		base = new(big.Int).SetBytes(getData(input, 0, baseLen))
		exp  = new(big.Int).SetBytes(getData(input, baseLen, expLen))
		mod  = new(big.Int).SetBytes(getData(input, baseLen+expLen, modLen))
	)
	if mod.BitLen() == 0 {
		// Modulo 0 is undefined, return zero
		return common.LeftPadBytes([]byte{}, int(modLen)), nil
	}
	return common.LeftPadBytes(base.Exp(base, exp, mod).Bytes(), int(modLen)), nil
}

var (
	// errNotOnCurve is returned if a point being unmarshalled as a bn256 elliptic
	// curve point is not on the curve.
	errNotOnCurve = errors.New("point not on elliptic curve")

	// errInvalidCurvePoint is returned if a point being unmarshalled as a bn256
	// elliptic curve point is invalid.
	errInvalidCurvePoint = errors.New("invalid elliptic curve point")
)

// newCurvePoint unmarshals a binary blob into a bn256 elliptic curve point,
// returning it, or an error if the point is invalid.
func newCurvePoint(blob []byte) (*bn256.G1, error) {
	p, onCurve := new(bn256.G1).Unmarshal(blob)
	if !onCurve {
		return nil, errNotOnCurve
	}
	gx, gy, _, _ := p.CurvePoints()
	if gx.Cmp(bn256.P) >= 0 || gy.Cmp(bn256.P) >= 0 {
		return nil, errInvalidCurvePoint
	}
	return p, nil
}

// newTwistPoint unmarshals a binary blob into a bn256 elliptic curve point,
// returning it, or an error if the point is invalid.
func newTwistPoint(blob []byte) (*bn256.G2, error) {
	p, onCurve := new(bn256.G2).Unmarshal(blob)
	if !onCurve {
		return nil, errNotOnCurve
	}
	x2, y2, _, _ := p.CurvePoints()
	if x2.Real().Cmp(bn256.P) >= 0 || x2.Imag().Cmp(bn256.P) >= 0 ||
		y2.Real().Cmp(bn256.P) >= 0 || y2.Imag().Cmp(bn256.P) >= 0 {
		return nil, errInvalidCurvePoint
	}
	return p, nil
}

// bn256Add implements a native elliptic curve point addition.
type bn256Add struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *bn256Add) RequiredGas(input []byte) uint64 {
	return config.Bn256AddGas
}

func (c *bn256Add) Run(input []byte) ([]byte, error) {
	log.Error("bn256Add", "run", common.Bytes2Hex(input))

	x, err := newCurvePoint(getData(input, 0, 64))
	if err != nil {
		log.Error("bn256Add", "err", err)
		return nil, err
	}
	y, err := newCurvePoint(getData(input, 64, 64))
	if err != nil {
		log.Error("bn256Add", "err", err)
		return nil, err
	}
	res := new(bn256.G1)
	res.Add(x, y)
	log.Error("bn256ScalarMul", "run", common.Bytes2Hex(res.Marshal()))

	return res.Marshal(), nil
}

// bn256ScalarMul implements a native elliptic curve scalar multiplication.
type bn256ScalarMul struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *bn256ScalarMul) RequiredGas(input []byte) uint64 {
	return config.Bn256ScalarMulGas
}

func (c *bn256ScalarMul) Run(input []byte) ([]byte, error) {
	log.Error("bn256ScalarMul", "run", common.Bytes2Hex(input))
	p, err := newCurvePoint(getData(input, 0, 64))
	if err != nil {
		return nil, err
	}
	res := new(bn256.G1)
	res.ScalarMult(p, new(big.Int).SetBytes(getData(input, 64, 32)))
	log.Error("bn256ScalarMul", "run", common.Bytes2Hex(res.Marshal()))

	return res.Marshal(), nil
}

var (
	// true32Byte is returned if the bn256 pairing check succeeds.
	true32Byte = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}

	// false32Byte is returned if the bn256 pairing check fails.
	false32Byte = make([]byte, 32)

	// errBadPairingInput is returned if the bn256 pairing input is invalid.
	errBadPairingInput = errors.New("bad elliptic curve pairing size")
)

// bn256Pairing implements a pairing pre-compile for the bn256 curve
type bn256Pairing struct{}

// RequiredGas returns the gas required to execute the pre-compiled contract.
func (c *bn256Pairing) RequiredGas(input []byte) uint64 {
	return config.Bn256PairingBaseGas + uint64(len(input)/192)*config.Bn256PairingPerPointGas
}

func (c *bn256Pairing) Run(input []byte) ([]byte, error) {
	// Handle some corner cases cheaply
	if len(input)%192 > 0 {
		return nil, errBadPairingInput
	}
	// Convert the input into a set of coordinates
	var (
		cs []*bn256.G1
		ts []*bn256.G2
	)
	for i := 0; i < len(input); i += 192 {
		c, err := newCurvePoint(input[i : i+64])
		if err != nil {
			return nil, err
		}
		t, err := newTwistPoint(input[i+64 : i+192])
		if err != nil {
			return nil, err
		}
		cs = append(cs, c)
		ts = append(ts, t)
	}
	// Execute the pairing checks and return the results
	if bn256.PairingCheck(cs, ts) {
		return true32Byte, nil
	}
	return false32Byte, nil
}
