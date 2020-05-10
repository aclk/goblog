module github.com/aclk/goblog/vipservice

go 1.13

replace github.com/aclk/goblog/common => /Users/arvin.he/go/src/github.com/aclk/goblog/common

require (
	github.com/aclk/goblog/common v0.0.0-20190723184755-e1fcaa3fae6e
	github.com/alexflint/go-arg v1.0.0
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/prometheus/client_golang v1.0.0
	github.com/sirupsen/logrus v1.4.2
	github.com/streadway/amqp v0.0.0-20190404075320-75d898a42a94
)
