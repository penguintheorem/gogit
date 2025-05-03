package util

import (
	"os"
	"testing"
)

func TestStoreObject(t *testing.T) {
	t.Run("creates_full_object_file_path", func(t *testing.T) {
		err := StoreObject("bc7774a7b18deb1d7bd0212d34246a9b1260ae17", "blob 12\x00"+"hello world!")
		if err != nil {
			t.Errorf("StoreObject failed to store the object with error %v", err)
		}

		objectFilePath := ".git/objects/" + "bc" + "/" + "7774a7b18deb1d7bd0212d34246a9b1260ae17"
		_, err = os.Stat(objectFilePath)
		if err != nil {
			t.Errorf("StoreObject failed to create the file path with error %v", err)
		}

		t.Cleanup(func() {
			defer os.RemoveAll(".git")
		})
	})

	t.Run("creates_non_empty_file_object", func(t *testing.T) {
		err := StoreObject("bc7774a7b18deb1d7bd0212d34246a9b1260ae17", "blob 12\x00"+"hello world!")
		if err != nil {
			t.Errorf("StoreObject failed to store the object with error %v", err)
		}

		objectFilePath := ".git/objects/" + "bc" + "/" + "7774a7b18deb1d7bd0212d34246a9b1260ae17"
		fileInfo, err := os.Stat(objectFilePath)
		if err != nil {
			t.Errorf("StoreObject failed to create the file path with error %v", err)
		}

		if fileInfo.Size() <= 0 {
			t.Errorf("The produced file size is 0")
		}

		t.Cleanup(func() {
			defer os.RemoveAll(".git")
		})
	})

}
