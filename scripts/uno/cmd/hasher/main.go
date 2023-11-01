package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	recursive                         bool
	doMD5, doSHA1, doSHA256, doSHA512 bool
)

func init() {
	rootCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "iterate the directory recursively")
	rootCmd.Flags().BoolVarP(&doMD5, "md5", "", false, "get the MD5 hash")
	rootCmd.Flags().BoolVarP(&doSHA1, "sha1", "", false, "get the SHA1 hash")
	rootCmd.Flags().BoolVarP(&doSHA256, "sha256", "", false, "get the SHA256 hash")
	rootCmd.Flags().BoolVarP(&doSHA512, "sha512", "", false, "get the SHA512 hash")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "hasher [file or directory]",
	Short: "generate an hash from a file and prints in json format",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		hashes := map[string]map[string]string{}
		file, err := os.Stat(args[0])
		if err != nil {
			return err
		}

		if file.IsDir() {
			if recursive {
				return filepath.Walk(args[0], func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}

					if info.IsDir() {
						return nil
					}

					hashes[path], err = hashFile(path)
					return err
				})
			} else {
				files, err := os.ReadDir(args[0])
				if err != nil {
					return err
				}

				for _, file := range files {
					if file.IsDir() {
						continue
					}

					hashes[file.Name()], err = hashFile(file.Name())
					if err != nil {
						return err
					}
				}
			}
		} else {
			hashes[args[0]], err = hashFile(args[0])
			if err != nil {
				return err
			}
		}

		jsonContent, err := json.MarshalIndent(hashes, "", "\t")
		if err != nil {
			return err
		}

		fmt.Println(string(jsonContent))
		return nil
	},
}

func hashFile(filepath string) (map[string]string, error) {
	var err error
	hashes := map[string]string{}

	if doMD5 {
		hashes["md5"], err = hashSum(md5.New(), filepath)
		if err != nil {
			return nil, err
		}
	}

	if doSHA1 {
		hashes["sha1"], err = hashSum(sha1.New(), filepath)
		if err != nil {
			return nil, err
		}
	}

	if doSHA256 {
		hashes["sha256"], err = hashSum(sha256.New(), filepath)
		if err != nil {
			return nil, err
		}
	}

	if doSHA512 {
		hashes["sha512"], err = hashSum(sha512.New(), filepath)
		if err != nil {
			return nil, err
		}
	}

	return hashes, nil
}

func hashSum(h hash.Hash, filepath string) (string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
