package main

//import (
//	"context"
//	"github.com/Postech-fiap-soat/ms-payment/internal/config"
//	"github.com/onsi/ginkgo/v2"
//	"github.com/onsi/gomega"
//	"testing"
//)
//
//func TestApplication(t *testing.T) {
//	gomega.RegisterFailHandler(ginkgo.Fail)
//	//junitReporter := reporters.NewJUnitReporter("../test/junit_report.xml")
//	ginkgo.RunSpecs(t, "location report application")
//}
//
//var _ = ginkgo.Describe("Soat payment microsservice", func() {
//	var (
//		cfg    *config.Config
//		cancel func()
//		ctx    context.Context
//	)
//	cfg = &config.Config{
//		ConnStr:                "mongodb://soatuser:soatpassword@localhost:27017/",
//		RabbitDialStr:          "amqp://soatuser:soatpassword@localhost:5672/",
//		RabbitExchange:         "amq.direct",
//		RabbitKey:              "soatkey",
//		MercadoPagoAccessToken: "test",
//		WebhookNotification:    "http://localhost/webhook",
//	}
//	ginkgo.BeforeEach(func() {
//		ctx, cancel = context.WithCancel(context.Background())
//		go func() {
//			LoadAPP(ctx, cfg)
//		}()
//
//	})
//	ginkgo.AfterEach(func() {
//		cancel()
//	})
//	ginkgo.Describe("description", func() {
//		ginkgo.When("when", func() {
//			ginkgo.It("it", func() {
//				//req, err := http.NewRequest("GET", "http://localhost:8001/hello", nil)
//				//gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
//				//res, err := http.DefaultClient.Do(req)
//				//gomega.Expect(err).ShouldNot(gomega.HaveOccurred())
//				//body, err := io.ReadAll(res.Body)
//				//gomega.Expect(string(body)).Should(gomega.Equal(1))
//			})
//		})
//	})
//})
