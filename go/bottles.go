package bottles

import "fmt"

func Verse(beers int) string {
  result := ""

  switch {
  case beers > 2:
    result = fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottles of beer on the wall.\n", beers, beers, beers - 1)
  case beers > 1:
    result = fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\nTake one down and pass it around, %d bottle of beer on the wall.\n", beers, beers, beers - 1)
  case beers > 0:
    result = fmt.Sprintf("%d bottle of beer on the wall, %d bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", beers, beers)
  default:
    result = fmt.Sprintf("No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n")
  }

  return result
}

func Verses(starting int, ending int) string {
  result := ""

  for beers := starting; beers >= ending; beers-- {
    result += Verse(beers)

    if beers != ending {
      result += "\n"
    }
  }

  return result
}

func Song() string {
  result := Verses(99, 0)

  return result
}
