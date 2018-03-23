class Bottles
  def song
    verses(99, 0)
  end

  def verses(start, finish)
    start.downto(finish).map { |beers| verse(beers) }.join("\n")
  end

  def verse(beers)
    case beers
    when 0
      "No more bottles of beer on the wall, no more bottles of beer.\n"\
      "Go to the store and buy some more, 99 bottles of beer on the wall.\n"
    when 1
      "#{beers} bottle of beer on the wall, #{beers} bottle of beer.\n"\
      "Take it down and pass it around, no more bottles of beer on the wall.\n"
    when 2
      "#{beers} bottles of beer on the wall, #{beers} bottles of beer.\n"\
      "Take one down and pass it around, #{beers-1} bottle of beer on the wall.\n"
    else
      "#{beers} bottles of beer on the wall, #{beers} bottles of beer.\n"\
      "Take one down and pass it around, #{beers-1} bottles of beer on the wall.\n"
    end
  end
end
