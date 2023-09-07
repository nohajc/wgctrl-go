//go:build !darwin && !windows

package wguser

func alt_sock_paths() ([]string, error) {
	return nil, nil
}
