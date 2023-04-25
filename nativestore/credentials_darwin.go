package nativestore

import "github.com/docker/docker-credential-helpers/osxkeychain"

var store = osxkeychain.Osxkeychain{}
