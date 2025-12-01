#!/usr/bin/env ruby

require 'colorize'

almanac = File.read(File.dirname(__FILE__) + "/input/day5_input.txt").split("\n\n")
# almanac = File.read("input/day5_sample.txt").split("\n\n")

### First, process the input into something useful

seedsVals = almanac[0].scan(/\d+/).map(&:to_i)
regexp = %r{(?<outStart>\d+) (?<inStart>\d+) (?<len>\d+)}

se2so = almanac[1].split("\n")[1..].map {
	|line| line.match(regexp).named_captures.transform_values!(&:to_i)
}

so2fe = almanac[2].split("\n")[1..].map {
	|line| line.match(regexp).named_captures.transform_values!(&:to_i)
}

fe2wa = almanac[3].split("\n")[1..].map {
	|line| line.match(regexp).named_captures.transform_values!(&:to_i)
}

wa2li = almanac[4].split("\n")[1..].map {
	|line| line.match(regexp).named_captures.transform_values!(&:to_i)
}

li2te = almanac[5].split("\n")[1..].map {
	|line| line.match(regexp).named_captures.transform_values!(&:to_i)
}

te2hu = almanac[6].split("\n")[1..].map {
	|line| line.match(regexp).named_captures.transform_values!(&:to_i)
}

hu2lo = almanac[7].split("\n")[1..].map {
	|line| line.match(regexp).named_captures.transform_values!(&:to_i)
}

puts "\nPart 1".red.on_black.underline

seedData = []

def mappedValue(input, map)
	map.each do |rangeData|
		range = (rangeData["inStart"]..rangeData["inStart"]+rangeData["len"]-1)
		if range.include?(input)
			transform = rangeData["outStart"] - rangeData["inStart"]
			return input + transform
		end
	end
	return input
end

seeds = []
seedsVals.each do |se|
	so = mappedValue(se, se2so)
	fe = mappedValue(so, so2fe)
	wa = mappedValue(fe, fe2wa)
	li = mappedValue(wa, wa2li)
	te = mappedValue(li, li2te)
	hu = mappedValue(te, te2hu)
	lo = mappedValue(hu, hu2lo)
	seeds << {:se => se, :so => so, :fe => fe, :wa => wa, :li => li, :te => te, :hu => hu, :lo => lo}
end

p seeds.map { |seed| seed[:lo] }.min

puts "\nPart 1".red.on_black.underline

seedData = []
