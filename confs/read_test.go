package confs_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"gitlab.com/byarbrough/wg-confs/confs"
)

func TestReadInterface(t *testing.T) {
	t.Parallel()

	t.Run("all values", func(t *testing.T) {
		buffer := []byte(`
	[Interface]
	PrivateKey = OPRGA+cLdEcPLIPKns/f1rqhzHTxuu1MS4ZNTxxinVc=
	ListenPort = 51820
	FwMark = 0x4D2
	`)

		// Convert buffer to section
		section, err := confs.GetSection(buffer, "Interface")
		if err != nil {
			t.Fatal(err)
		}
		// Get the configuration as a Config
		got, err := confs.ReadInterface(section)
		if err != nil {
			t.Error(err)
		}

		// Expected values
		privateKey := "OPRGA+cLdEcPLIPKns/f1rqhzHTxuu1MS4ZNTxxinVc="

		want := &confs.Config{
			PrivateKey: privateKey,
			ListenPort: uint16(51820),
			FwMark:     uint32(0x4D2),
			Peers:      nil,
		}

		if !cmp.Equal(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})

	t.Run("required values", func(t *testing.T) {
		buffer := []byte(`
		[Interface]
		PrivateKey = kFDRIITet2LwACfznYaMdx4YlXsIlKdqcnuFE1zY32Y=
		`)

		// Convert buffer to section
		section, err := confs.GetSection(buffer, "Interface")
		if err != nil {
			t.Fatal(err)
		}
		// Get the configuration as a Config
		got, err := confs.ReadInterface(section)
		if err != nil {
			t.Error(err)
		}

		// Expected values
		privateKey := "kFDRIITet2LwACfznYaMdx4YlXsIlKdqcnuFE1zY32Y="

		want := new(confs.Config)
		want.PrivateKey = privateKey

		if !cmp.Equal(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})
}

func TestReadPeer(t *testing.T) {
	t.Parallel()

	t.Run("required values", func(t *testing.T) {
		buffer := []byte(`
		[Peer]
		PublicKey = 4F5mIj+fcdE4hTEYLnjJHls+Zigy++wy5yiS6B9k8kM=
		`)

		// Convert buffer to section
		section, err := confs.GetSection(buffer, "Peer")
		if err != nil {
			t.Fatal(err)
		}
		// Get the configuration as a Config
		got, err := confs.ReadPeer(section)
		if err != nil {
			t.Error(err)
		}

		// Expected values
		publicKey := "4F5mIj+fcdE4hTEYLnjJHls+Zigy++wy5yiS6B9k8kM="

		want := &confs.Peer{
			PublicKey: publicKey,
		}
		if !cmp.Equal(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})
}

func TestReadConfig(t *testing.T) {
	t.Parallel()

	t.Run("interface only", func(t *testing.T) {
		buffer := []byte(`
			[Interface]
			PrivateKey = aClVSMm9VEDx3aYAXg4FYKhAvchXw10e0IABJgrBjUM=
			`)

		// Get the configuration as a Config
		got, err := confs.ReadConfig(buffer)
		if err != nil {
			t.Error(err)
		}

		// Expected values
		privateKey := "aClVSMm9VEDx3aYAXg4FYKhAvchXw10e0IABJgrBjUM="

		want := new(confs.Config)
		want.PrivateKey = privateKey

		if !cmp.Equal(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})

	t.Run("one peer", func(t *testing.T) {
		buffer := []byte(`
			[Interface]
			PrivateKey = 6CajzW/qXuB07um+50CYU+N4Tucud3xtllh6JYsCEUQ=

			[Peer]
			PublicKey =3VpNW2azh4M61+ziZX0O768l2IemS5QhACgQMaMfIFs=
			`)

		// Get the configuration as a Config
		got, err := confs.ReadConfig(buffer)
		if err != nil {
			t.Error(err)
		}

		// Expected values
		privateKey := "6CajzW/qXuB07um+50CYU+N4Tucud3xtllh6JYsCEUQ="
		peerPublicKey := "3VpNW2azh4M61+ziZX0O768l2IemS5QhACgQMaMfIFs="
		newPeer := &confs.Peer{
			PublicKey: peerPublicKey,
		}
		want := &confs.Config{
			PrivateKey: privateKey,
		}
		want.Peers = append(want.Peers, newPeer)

		if !cmp.Equal(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})

	t.Run("many peers", func(t *testing.T) {
		buffer := []byte(`
		[Interface]
		PrivateKey = 6CajzW/qXuB07um+50CYU+N4Tucud3xtllh6JYsCEUQ=

		[Peer]
		PublicKey =3VpNW2azh4M61+ziZX0O768l2IemS5QhACgQMaMfIFs=

		[Peer]
		PublicKey = o2voeRt/89DwDbB38oiZ92PeGZb30/jdQdQnLECPPDE=
		`)

		// Get the configuration as a Config
		got, err := confs.ReadConfig(buffer)
		if err != nil {
			t.Error(err)
		}

		// Expected values
		privateKey := "6CajzW/qXuB07um+50CYU+N4Tucud3xtllh6JYsCEUQ="
		peerPublicKey := "3VpNW2azh4M61+ziZX0O768l2IemS5QhACgQMaMfIFs="
		peer1 := &confs.Peer{
			PublicKey: peerPublicKey,
		}
		peerPublicKey = "o2voeRt/89DwDbB38oiZ92PeGZb30/jdQdQnLECPPDE="
		peer2 := &confs.Peer{
			PublicKey: peerPublicKey,
		}

		want := &confs.Config{
			PrivateKey: privateKey,
		}
		want.Peers = append(want.Peers, peer1, peer2)

		if !cmp.Equal(got, want) {
			t.Fatalf("got %v want %v", got, want)
		}
	})
}
