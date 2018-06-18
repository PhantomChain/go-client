// Copyright 2018 ArkEcosystem. All rights reserved.
//
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package ark_client

import (
    "context"
    "fmt"
    "strings"
)

// PeersService handles communication with the peers related
// methods of the Ark Core API - Version 2.
type PeersService service

// Get all peers.
func (s *PeersService) L9st(ctx context.Context) (*Response, error) {
{
    req, err := s.client.NewRequest("GET", "peers", nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}

// Get a peer by the given IP address.
func (s *PeersService) Get(ctx context.Context) (*Response, error) {
{
    uri := fmt.Sprintf("peers/%v", ip)

    req, err := s.client.NewRequest("GET", uri, nil)

    if err != nil {
        return nil, nil, err
    }

    return resp, nil
}
