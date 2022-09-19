// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rocketmq

import (
	"errors"
	"github.com/apache/incubator-eventmesh/eventmesh-server-go/plugin"
	"github.com/apache/incubator-eventmesh/eventmesh-server-go/plugin/connector"
)

func init() {
	plugin.Register("rocketmq", &ProducerFactory{})
	plugin.Register("rocketmq", &ConsumerFactory{})
	plugin.Register("rocketmq", &ResourceFactory{})
}

type ProducerFactory struct {
	plugin.Plugin
	properties map[string]string
}

func (f *ProducerFactory) Type() string {
	return connector.ProducerPluginType
}

func (f *ProducerFactory) Setup(name string, dec plugin.Decoder) error {
	if dec == nil {
		return errors.New(" producer config decoder empty")
	}
	properties := make(map[string]string)
	if err := dec.Decode(properties); err != nil {
		return err
	}
	f.properties = properties
	return nil
}

func (f *ProducerFactory) Get() (connector.Producer, error) {
	producer := NewProducer()
	err := producer.InitProducer(f.properties)
	if err != nil {
		return nil, err
	}
	err = producer.Start()
	if err != nil {
		return nil, err
	}
	return producer, nil
}

type ConsumerFactory struct {
	plugin.Plugin
	properties map[string]string
}

func (c *ConsumerFactory) Type() string {
	return connector.ConsumerPluginType
}

func (c *ConsumerFactory) Setup(name string, dec plugin.Decoder) error {
	if dec == nil {
		return errors.New(" producer config decoder empty")
	}
	properties := make(map[string]string)
	if err := dec.Decode(properties); err != nil {
		return err
	}
	c.properties = properties
	return nil
}

func (c *ConsumerFactory) Get() (connector.Consumer, error) {
	consumer := NewConsumer()
	err := consumer.InitConsumer(c.properties)
	if err != nil {
		return nil, err
	}
	err = consumer.Start()
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

type ResourceFactory struct {
	plugin.Plugin
}

func (f *ResourceFactory) Type() string {
	return connector.ResourcePluginType
}

func (f *ResourceFactory) Setup(name string, dec plugin.Decoder) error {
	return nil
}

func (f *ResourceFactory) Get() (connector.Resource, error) {
	return &Resource{}, nil
}