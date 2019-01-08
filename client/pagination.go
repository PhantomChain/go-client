// This file is part of PHATOM Go Client.
//
// (c) PhantomChain <info@phantom.org>
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package client

type Pagination struct {
	Page  int `url:"page"`
	Limit int `url:"limit"`
}
