package search

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pkg/errors"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 배열은 정렬이 되어 있는 것으로 간주

	idx, err := binarySearch(arr, 3, 0, len(arr)-1)
	require.NoError(t, err)
	require.Equal(t, 2, idx)

	idx, err = binarySearch(arr, 8, 0, len(arr)-1)
	require.NoError(t, err)
	require.Equal(t, 7, idx)

	_, err = binarySearch(arr, -8, 0, len(arr)-1)
	require.Error(t, err)
}

func binarySearch[T int](arr []T, value T, low, high int) (index int, err error) {
	if length := len(arr); length == 0 || low > high {
		return -1, errors.New("not exist")
	} else if mid := (high + low) / 2; arr[mid] == value {
		return mid, nil
	} else if arr[mid] > value {
		return binarySearch(arr, value, low, mid-1)
	} else {
		return binarySearch(arr, value, mid+1, high)
	}
}
