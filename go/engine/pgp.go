package engine

//
// engine.PGPEngine is a class for optionally generating PGP keys,
// and pushing them into the keybase sigchain via the Delegator.
//

import (
	stderrors "errors"
	triplesec "github.com/keybase/go-triplesec"
	"github.com/keybase/client/go/libkb"
)

type PGPEngine struct {
	me     *libkb.User
	bundle *libkb.PgpKeyBundle
	arg    PGPEngineArg
	epk    string
}

type PGPEngineArg struct {
	Gen        *libkb.PGPGenArg
	Pregen     *libkb.PgpKeyBundle
	SigningKey libkb.GenericKey
	NoSave     bool
	PushSecret bool
}

func (s *PGPEngine) loadMe() (err error) {
	s.me, err = libkb.LoadMe(libkb.LoadUserArg{PublicKeyOptional: true})
	return err
}

func (s *PGPEngine) generateKey(ctx *Context) (err error) {
	gen := s.arg.Gen
	if err = gen.CreatePgpIDs(); err != nil {
		return
	}
	s.bundle, err = libkb.NewPgpKeyBundle(*gen, ctx.LogUI)
	return
}

func (s *PGPEngine) saveLKS(ctx *Context) (err error) {
	var lks *libkb.LKSec
	if lks, err = libkb.NewLKSForEncrypt(ctx.SecretUI); err != nil {
		return
	}
	_, err = libkb.WriteLksSKBToKeyring(s.me.GetName(), s.bundle, lks, ctx.LogUI)
	return
}

var ErrKeyGenArgNoDefNoCustom = stderrors.New("invalid args:  NoDefPGPUid set, but no custom PGPUids.")

func NewPGPEngine(arg PGPEngineArg) *PGPEngine {
	return &PGPEngine{arg: arg}
}

func (s *PGPEngine) Name() string {
	return "PGP"
}

func (e *PGPEngine) GetPrereqs() EnginePrereqs {
	return EnginePrereqs{
		Session: true,
	}
}

func (k *PGPEngine) RequiredUIs() []libkb.UIKind {
	return []libkb.UIKind{
		libkb.LogUIKind,
		libkb.SecretUIKind,
	}
}

func (s *PGPEngine) SubConsumers() []libkb.UIConsumer {
	return nil
}

func (s *PGPEngine) init() (err error) {
	if s.arg.Gen != nil {
		err = s.arg.Gen.Init()
	}
	return err
}

func (s *PGPEngine) Run(ctx *Context, args interface{}, reply interface{}) (err error) {
	G.Log.Debug("+ PGPEngine::Run")
	defer func() {
		G.Log.Debug("- PGPEngine::Run -> %s", libkb.ErrToOk(err))
	}()

	if err = s.init(); err != nil {
	} else if err = s.loadMe(); err != nil {
	} else if err = s.loadSigningKey(ctx); err != nil {
	} else if err = s.generate(ctx); err != nil {
	} else {
		err = s.push(ctx)
	}

	return
}

func (s *PGPEngine) loadSigningKey(ctx *Context) (err error) {
	if s.arg.SigningKey != nil {
		G.Log.Debug("| PGPEngine: argSigning Key was passed in")
		return
	}

	if !s.me.HasActiveKey() {
		G.Log.Debug("| PGPEngine: no active key found, so assuming set of eldest key")
		return
	}

	s.arg.SigningKey, err = G.Keyrings.GetSecretKey(libkb.SecretKeyArg{
		All:    true,
		Me:     s.me,
		Ui:     ctx.SecretUI,
		Reason: "sign new key",
	})

	return err
}

func (s *PGPEngine) generate(ctx *Context) (err error) {

	G.Log.Debug("+ PGP::Generate")
	defer func() {
		G.Log.Debug("- PGP::Generate -> %s", libkb.ErrToOk(err))
	}()

	G.Log.Debug("| GenerateKey")
	if s.arg.Pregen != nil {
		s.bundle = s.arg.Pregen
	} else if s.arg.Gen == nil {
		err = libkb.InternalError{"PGPEngine: need either Gen or Pregen"}
		return
	} else if err = s.generateKey(ctx); err != nil {
		return
	}

	G.Log.Debug("| WriteKey")
	if s.arg.NoSave {
	} else if err = s.saveLKS(ctx); err != nil {
		return
	}

	if !s.arg.PushSecret {
	} else if err = s.prepareSecretPush(ctx); err != nil {
		return
	}
	return

}

func (s *PGPEngine) prepareSecretPush(ctx *Context) (err error) {
	var tsec *triplesec.Cipher
	var skb *libkb.SKB
	if tsec, err = G.LoginState.GetVerifiedTriplesec(ctx.SecretUI); err != nil {
	} else if skb, err = s.bundle.ToSKB(tsec); err != nil {
	} else {
		s.epk, err = skb.ArmoredEncode()
	}
	return
}

func (s *PGPEngine) push(ctx *Context) (err error) {
	G.Log.Debug("+ PGP::Push")
	defer func() {
		G.Log.Debug("- PGP::Push -> %s", libkb.ErrToOk(err))
	}()
	d := libkb.Delegator{
		ExistingKey:       s.arg.SigningKey,
		NewKey:            s.bundle,
		Me:                s.me,
		Expire:            libkb.KEY_EXPIRE_IN,
		Sibkey:            true,
		EncodedPrivateKey: s.epk,
	}
	err = d.Run()
	return
}
