package geometry

import "errors"

// Retorna el cubo de un entero, de lo contrario un error
func CubeVolume(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	} else {
		return 0, errors.New("Zero length edge is not allowed")
	}
}
