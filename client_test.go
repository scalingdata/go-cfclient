package cfclient

import (
	"testing"

	"github.com/onsi/gomega"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDefaultConfig(t *testing.T) {
	Convey("Default config", t, func() {
		c := DefaultConfig()
		So(c.ApiAddress, ShouldEqual, "http://api.bosh-lite.com")
		So(c.Username, ShouldEqual, "admin")
		So(c.Password, ShouldEqual, "admin")
		So(c.SkipSslValidation, ShouldEqual, false)
		So(c.Token, ShouldEqual, "")
	})
}

func TestMakeRequest(t *testing.T) {
	Convey("Test making request", t, func() {
		setup(MockRoute{"GET", "/v2/organizations", listOrgsPayload})
		defer teardown()
		c := &Config{
			ApiAddress:        server.URL,
			Username:          "foo",
			Password:          "bar",
			SkipSslValidation: true,
		}
		client, err := NewClient(c)
		So(err, ShouldBeNil)
		req := client.NewRequest("GET", "/v2/foobar")
		resp, err := client.DoRequest(req)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
	})
}

func TestCustomCACert(t *testing.T) {
	Convey("Test making request with user-specified CA cert", t, func() {
		setup(MockRoute{"GET", "/v2/organizations", listOrgsPayload})
		defer teardown()
		c := &Config{
			ApiAddress: server.URL,
			Username:   "foo",
			Password:   "bar",
		}
		c.CAPem = `
-----BEGIN CERTIFICATE-----
MIIFBTCCAu2gAwIBAgIJAKkiyTwYVnGEMA0GCSqGSIb3DQEBBQUAMBkxFzAVBgNV
BAMMDmNlcnQtYXV0aG9yaXR5MB4XDTE2MTIwNTAzMzI1OVoXDTE3MTIwNTAzMzI1
OVowGTEXMBUGA1UEAwwOY2VydC1hdXRob3JpdHkwggIiMA0GCSqGSIb3DQEBAQUA
A4ICDwAwggIKAoICAQDPDtUfab28NUUrwL99P0qUTXsqGWV7VOvNHmj0SFcwMBvM
66COR+7aDV0w+I8vCP0EJaFvekQ8vVd2XelO2svpAZ8xtMMzBk6C7eh0QgXoPkKF
deBpgK0htyp93xiLnt+kKWpWR0x0hRm1JbrXRw7HjT1VSKUmDoQsPadV+Z4psPgi
yHVO07ZFQl21PAZJMh7TyeSHcDrGTlRe9G7WGvOoF8Nlx0kZwOILIzjjjx7cbA94
GO8qdQRQVc84ZMtuD7oRGof69wHhLJjxdJgDhw3Dpvdvw1bmcsuzftq4wR+CjHp1
ykPh6HyD669Sog7RwSAu6jkxQZBhpWXbijnN1oB7KNy0+jFmLFPYj1r9Q0clhDtF
ZboYcolL5rgR3WIu7peEJXptLKjRk11i74RVv+TDBJXTQBdfG/ea/m1SqgzhB01N
pvHk0DFcs+kEG5Q2gj/AL/5kHEr4ZshjadD46/CxhTUAYU1YDjpMeO9kxcgK/544
wUgkVf9hBxVyl/nRW8fj7uOjFefW/AEoq+z9oh5vI2ttGoT9boz3apQbbzSkw7Jd
5RdEmIAdGgUKOzqcgeMEZB9b51Xpd4WNz9TE7nw1rLo6bzo3eupFT5CYyVk4cQrE
h0Yv/GHbLxUGqUgNyjJrEBceBXRd5dTDdGygjKGkZr145ahfhj5U7umJNKLflwID
AQABo1AwTjAdBgNVHQ4EFgQUaDbOYrRmNuIJVeERgl5zSerq9p4wHwYDVR0jBBgw
FoAUaDbOYrRmNuIJVeERgl5zSerq9p4wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0B
AQUFAAOCAgEAALFbQvN4vsP3h/YyJAYDK8VKigM/c0yfYA0ObgNeBA0giUZ7aLDk
HFIE8Z7nuds7guTJyv5MpDbHd8VBbtW3ujA3wq6DdxxpNqAI41rUZK20whyhwz8X
zLnHcU07aMusZp5lCbSyIMU/K5RlflAxABWiePS4wf0aOcb0VoNSxgiKVYrN8p3N
piUJalCW4swif1ws1G6P3qb8u1bS7qQAWB+kGKBBDDA/qrqUb9COyVqCzhDLOZSi
oarqrC2voq+WErxUnkygdcfpyEWxmvbndh/Lesh6seSOB+mvvxWJJKjX+Fi5vBjf
9E41n+p62M34KSfvrTlQ8ewLM0o0XPIc7Xrnoe30mJMhOrVJ0whCa0e/Jd4PJzu5
wiQJUliCBpj3pFzDpy+rB2EP/bGiPlJkFFd95peqQxfb+CxysPv6b6wZV3nGQZ0G
LZw7t1HA6fCAu1CRzDzE8hx6ao1xXmSkdf5RocwLk7M8vVw1AxfdiB91/P+QNn3d
6M6oSvaHcqpXWPFk+HxCgTubgBgE1Qskwl9AlmCDsH2q5D7CIOAOJaNWN3vmC5dl
uD8vi8TOoPyx7Wm53ju2nRYoOGHiTBiJra/u2AEqWduz/gjBH3fHe5VOOU15/v72
94fWGMFDhYFNFfTWLr1iOyZSP8DpAxRSJkRvpCQIxubZXoq3Qki4flg=
-----END CERTIFICATE-----
`
		client, err := NewClient(c)
		So(err, ShouldBeNil)
		req := client.NewRequest("GET", "/v2/foobar")
		resp, err := client.DoRequest(req)
		So(err, ShouldBeNil)
		So(resp, ShouldNotBeNil)
	})
}

func TestTokenRefresh(t *testing.T) {
	gomega.RegisterTestingT(t)
	Convey("Test making request", t, func() {
		setup(MockRoute{"GET", "/v2/organizations", listOrgsPayload})
		c := &Config{
			ApiAddress: server.URL,
			Username:   "foo",
			Password:   "bar",
		}
		client, err := NewClient(c)
		So(err, ShouldBeNil)

		token, err := client.GetToken()
		So(err, ShouldBeNil)

		gomega.Consistently(token).Should(gomega.Equal("bearer foobar2"))
		// gomega.Eventually(client.GetToken(), "3s").Should(gomega.Equal("bearer foobar3"))
	})
}
