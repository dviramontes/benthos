package integration

import (
	"fmt"
	"testing"
	"time"

	"github.com/ory/dockertest/v3"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var _ = registerIntegrationTest("amqp_0_9", func(t *testing.T) {
	t.Parallel()

	pool, err := dockertest.NewPool("")
	require.NoError(t, err)

	pool.MaxWait = time.Second * 30

	resource, err := pool.Run("rabbitmq", "latest", nil)
	require.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, pool.Purge(resource))
	})

	resource.Expire(900)
	require.NoError(t, pool.Retry(func() error {
		client, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@localhost:%v/", resource.GetPort("5672/tcp")))
		if err == nil {
			client.Close()
		}
		return err
	}))

	template := `
output:
  amqp_0_9:
    url: amqp://guest:guest@localhost:$PORT/
    max_in_flight: $MAX_IN_FLIGHT
    exchange: exchange-$ID
    key: benthos-key
    exchange_declare:
      enabled: true
      type: direct
      durable: true
    metadata:
      exclude_prefixes: [ $OUTPUT_META_EXCLUDE_PREFIX ]

input:
  amqp_0_9:
    url: amqp://guest:guest@localhost:$PORT/
    auto_ack: $VAR1
    queue: queue-$ID
    queue_declare:
      durable: true
      enabled: true
    bindings_declare:
      - exchange: exchange-$ID
        key: benthos-key
`
	suite := integrationTests(
		integrationTestOpenClose(),
		integrationTestMetadata(),
		integrationTestMetadataFilter(),
		integrationTestSendBatch(10),
		integrationTestStreamSequential(1000),
		integrationTestStreamParallel(1000),
		integrationTestStreamParallelLossy(1000),
		integrationTestStreamParallelLossyThroughReconnect(1000),
	)
	suite.Run(
		t, template,
		testOptSleepAfterInput(500*time.Millisecond),
		testOptSleepAfterOutput(500*time.Millisecond),
		testOptPort(resource.GetPort("5672/tcp")),
		testOptVarOne("false"),
	)
	t.Run("with max in flight", func(t *testing.T) {
		t.Parallel()
		suite.Run(
			t, template,
			testOptSleepAfterInput(500*time.Millisecond),
			testOptSleepAfterOutput(500*time.Millisecond),
			testOptPort(resource.GetPort("5672/tcp")),
			testOptVarOne("false"),
			testOptMaxInFlight(10),
		)
	})
	t.Run("with auto ack", func(t *testing.T) {
		t.Parallel()
		integrationTests(
			integrationTestOpenClose(),
			integrationTestMetadata(),
			integrationTestSendBatch(10),
			integrationTestStreamSequential(100),
			integrationTestStreamParallel(100),
		).Run(
			t, template,
			testOptVarOne("true"),
			testOptSleepAfterInput(100*time.Millisecond),
			testOptSleepAfterOutput(100*time.Millisecond),
			testOptPort(resource.GetPort("5672/tcp")),
		)
	})
})
