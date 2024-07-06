package nats_consumer

let tagNATSServer = "v2.11.0-preview.1-alpine3.19"
let tagNATSBox = "0.14.3"
let tagBenthos = "4"

#NATSClusterOpts = [
	"--cluster", "nats://0.0.0.0:6222",
	"--cluster-name", "test-cluster",
]

#NATS: {
	image: "synadia/nats-server:\(tagNATSServer)"
	entrypoint: *["nats-server", "-js", #NATSClusterOpts...] | [...string]

	depends_on?: [...string]
	ports?: [...string]
}

services: {
	nats1: #NATS
	nats2: #NATS & {
		entrypoint: ["nats-server", "-js", #NATSClusterOpts..., ]
	}
	nats3: #NATS & {
	}

	box: {
		image:       "natsio/nats-box:\(tagNATSBox)"
		command:     "sleep infinity"
		working_dir: "/work"
		volumes: ["./backup:/work/backup:ro"]
		environment: NATS_URL: "nats://nats:4222/"
	}

	benthos: {
		image:       "jeffail/benthos:\(tagBenthos)"
		working_dir: "/work"
		environment: NATS_URL: "nats://nats1:4222/"
		volumes: ["./seed.yaml:/benthos.yaml:ro"]
	}
}
