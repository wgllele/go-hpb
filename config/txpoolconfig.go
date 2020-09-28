// Copyright 2018 The go-hpb Authors
// Modified based on go-ethereum, which Copyright (C) 2014 The go-ethereum Authors.
//
// The go-hpb is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-hpb is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-hpb. If not, see <http://www.gnu.org/licenses/>.

package config

import "time"

// TxPoolConfiguration txpool config.
type TxPoolConfiguration struct {
	PriceLimit   uint64        // Minimum gas price to enforce for acceptance into the pool
	PriceBump    uint64        // Minimum price bump percentage to replace an already existing transaction (nonce)
	AccountSlots uint64        // Minimum number of executable transaction slots guaranteed per account
	GlobalSlots  uint64        // Maximum number of executable transaction slots for all accounts
	AccountQueue uint64        // Maximum number of non-executable transaction slots permitted per account
	GlobalQueue  uint64        // Maximum number of non-executable transaction slots for all accounts
	Lifetime     time.Duration // Maximum amount of time non-executable transaction are queued
}

// DefaultTxPoolConfig default txpool config.
var DefaultTxPoolConfig = TxPoolConfiguration{
	PriceLimit:   1,
	PriceBump:    10,
	AccountSlots: 60000,
	GlobalSlots:  1000000,
	AccountQueue: 30000,
	GlobalQueue:  1000000,
	Lifetime:     30 * time.Minute,
}
