/*
 * 指定ディレクトリにあるテキストを Open JTalk に渡して wav ファイルを作り、それを aplay で再生します
 */
package main

import (
  "fmt"
  "os"
  "flag"
  "sort"
  "io/ioutil"
  "os/exec"
)

func main() {
  input := flag.String("input", "source/", "input text dir")
  tmpwav := flag.String("tmpwav", "output.wav", "temporary wav file")
  jtalk := flag.String("open_jtalk", "/usr/bin/open_jtalk", "Open JTalk")
  dic := flag.String("x", "/var/lib/mecab/dic/open-jtalk/naist-jdic/", "dictionary for Open JTalk")
  voice := flag.String("m", "/usr/share/hts-voice/mei/mei_normal.htsvoice", "voice for Open JTalk")
  aplay := flag.String("aplay", "aplay", "wav player")
  flag.Parse()


  list, err := ioutil.ReadDir(*input)
  if err != nil {
    fmt.Fprintf(os.Stderr, "%v", err)
    os.Exit(1)
  }

  var files []string
  for _, finfo := range list {
    if finfo.IsDir() { 
      continue
    }
    files = append(files, finfo.Name())
  }
  sort.Strings(files)

  for _, file := range files {
    f := *input + file
    // open jtalk
    cmd := exec.Command(*jtalk)
    cmd.Args = []string {
      "-x", *dic,
      "-m", *voice,
      f,
      "-ow", *tmpwav,
    }
    cmd.Run()

    // aplay
    cmd = exec.Command(*aplay, *tmpwav)
    cmd.Run()
  }
}
