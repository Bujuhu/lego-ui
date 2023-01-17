package config

import(	
	"strings"
	"time"
	"github.com/go-acme/lego/v4/registration"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fyne.io/fyne/v2"
	"log")

// The Config struct contains parameters that can be customized by the User through the Settings Page
type Config struct {

	preferences fyne.Preferences
	//Is only set, if a user is registered with LetsEncrypt
	Registration *registration.Resource
	
	//Private Key used for certificates
	Key          crypto.PrivateKey

	//After how many seconds, should we cancel the DNS verification request
	PropagationTimeout time.Duration

	//What interval should we check DNS Servers, if the verification succeeded
	PollingIntervall time.Duration
}

type CertificateDefinition struct {
	id string
	domain string
	dnsProvider string
}

func New(pref fyne.Preferences) Config {
	return Config { preferences: pref }
}

// Randomly generates a new private Key, which will be used for future certificates
func NewPrivateKey() crypto.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	return privateKey
}

func (u Config) GetRegistration() *registration.Resource {
	return u.Registration
}
func (u *Config) GetPrivateKey() crypto.PrivateKey {
	return u.Key
}

// E-Mail Adress needed to regstier with LetsEncrypt
func (u *Config) GetEmail() string {
	return u.preferences.StringWithFallback("user.email", "")
}

// E-Mail Adress needed to regstier with LetsEncrypt
func (u *Config) SetEmail(email string) {
	u.preferences.SetString("user.email", email)
}

func (u *Config) GetTOSAgree() bool {
	return u.preferences.BoolWithFallback("user.TOSAgree", false)
}

func (u *Config) SetTOSAgree(tosAgree bool) {
	u.preferences.SetBool("user.TOSAgree", tosAgree)
}

func (u *Config) GetCertificateDefinitions() []CertificateDefinition {
	certificateDefintions:= []CertificateDefinition{}
	serialisedDefinitions := u.preferences.StringWithFallback("user.certificateDefinitions", "")
	definitions := strings.Split(serialisedDefinitions, ",")
	for _, sd := range definitions {
		elements := strings.Split(sd, ":")
		certificateDefintions = append(certificateDefintions, CertificateDefinition {
			id: elements[0],
			domain: elements[1],
			dnsProvider: elements[2],
		})
    }
	return certificateDefintions
}

func (u *Config) SetCertificateDefinitions(definitions []CertificateDefinition) {
	serialisedDefinitions := []string{}
	for _, d := range definitions  {
		serialisedDefinitions = append(serialisedDefinitions, d.id + ":" + d.domain + ":" + d.dnsProvider)
	}

	u.preferences.SetString("user.certificateDefinitions", strings.Join(serialisedDefinitions, ","))
}