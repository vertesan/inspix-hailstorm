package story

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"vertesan/hailstorm/master"
	"vertesan/hailstorm/rich"
	"vertesan/hailstorm/utils"
)

func ParseStoryText() {

	advMaster := []*master.AdvDatas{}

	if err := utils.ReadFromYamlFile("masterdata/AdvDatas.yaml", &advMaster); err != nil {
		rich.PanicError("failed to read yaml master file", err)
	}

	advMap := make(map[string]*master.AdvDatas, len(advMaster))

	for _, adv := range advMaster {
		advMap[fmt.Sprintf("story_main_%d.txt", adv.ScriptId)] = adv
	}

	err := filepath.WalkDir("./cache/txt", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(path) == ".txt" {
			rich.Info("parsing story file: %q", path)

			fp, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fp.Close()

			adv := advMap[filepath.Base(path)]
			os.MkdirAll("cache/story", 0755)
			fw, err := os.Create(fmt.Sprintf("cache/story/%d %s %s.txt", adv.ScriptId, sanitizeFilename(adv.Name), sanitizeFilename(adv.Description)))
			if err != nil {
				return err
			}
			defer fw.Close()
			bufw := bufio.NewWriter(fw)

			rCapture := regexp.MustCompile(`^\s*\[メッセージ表示\s+(\p{L}+)\s+\w+@\w+\s+([\s\S]+?)\]\s*$`)
			rNovel := regexp.MustCompile(`^ *\[ノベルテキスト追加\s+([\s\S]+?)\s*[\w@]*\] *$`)
			rNovelRm := regexp.MustCompile(`^ *\[ノベルテキスト削除`)

			scanner := bufio.NewScanner(fp)
			for scanner.Scan() {
				line := scanner.Text()

				matches := rCapture.FindStringSubmatch(line)
				if len(matches) >= 3 {
					chara := matches[1]
					dialog := matches[2]

					bufw.WriteString(chara)
					bufw.WriteString("\n")
					bufw.WriteString(strings.ReplaceAll(strings.ReplaceAll(dialog, "[r]", "\n"), "[Space]", " "))
					bufw.WriteString("\n\n")
					continue
				}

				matches = rNovel.FindStringSubmatch(line)
				if len(matches) >= 2 {
					novelText := matches[1]
					bufw.WriteString(strings.ReplaceAll(strings.ReplaceAll(novelText, "[r]", "\n"), "[Space]", " "))
					bufw.WriteString("\n")
					continue
				}

				exist := rNovelRm.MatchString(line)
				if exist {
					bufw.WriteString("\n")
					continue
				}

			}

			bufw.Flush()
		}

		return nil
	})

	if err != nil {
		rich.PanicError("failed to parse story text: ", err)
	}
}

func sanitizeFilename(s string) string {
	invalid := []rune{'<', '>', ':', '"', '/', '\\', '|', '?', '*'}
	return strings.Map(func(r rune) rune {
		for _, inv := range invalid {
			if r == inv {
				return -1
			}
		}
		return r
	}, s)
}
