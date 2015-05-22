# Cisco Micro Services

[![Build Status](https://drone.io/github.com/triforkse/cisco-micro/status.png)](https://drone.io/github.com/triforkse/cisco-micro/latest)
[![Coverage Status](https://coveralls.io/repos/triforkse/cisco-micro/badge.svg?branch=master)](https://coveralls.io/r/triforkse/cisco-micro?branch=master)
# Installation

clone this repo. run:

```bash
make setup
make build
```

# Usage

## Creating a cluster

First you will have to write a configuration file for your cluster. You can
create a cluster on the `Amazon Cloud` or `Google Cloud`. You can start by
copying an example configuration from the `samples/` directory.

Or by running:

```bash
build/micro -provider=aws init
```

You can use `aws` or `gce` as cloud providers. The resulting configuration file
will be stored in `infrastructure.json`

Here is how an example amazon setup could look:

```json
{
  "id": "your-aws-1",
  "provider": "aws",
  "secret_key": "af4YY1yeXeQkByTYUYFtBBLUjL4YXXFTHaFBvaDb",
  "access_key": "7AIAIOEIEEEAZQ4AIE67",
  "region": "eu-west-1"
}
```

The `id` is the name of your cluster, you should make this unique.
You can have several cluster with the same cloud provider as long as they
use different ids.

The `provider` property determines what cloud provider you are going to use,
available values are `aws` and `gce`.

The `properties` are provider specific. Look at the example to see what is
required. For instance for the google cloud provider you will need to download
an OAuth profile and add the values from the download json file to
the properties in your configuration.

Once your configuration file is done, you can run:

```bash
./build/micro -config=<path to config file>
```

This will create your cluster and produce some state files in the `.micro`
directory. These you will want to check in to your CVS.

## Updating your cluster

If you run the command again after altering the configuration file, the cluster
will be updated to reflect the changes.
