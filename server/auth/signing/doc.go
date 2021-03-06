// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Package signing provides interfaces to sign arbitrary small blobs with
// RSA-SHA256 signature (PKCS1v15) and verify such signatures.
//
// Each service has its own private keys it uses for signing, with public
// certificates served over HTTPS. Other services may use the public keys
// to authenticate data generated by the service. It is useful, for example, for
// authenticating PubSub messages payload.
package signing
