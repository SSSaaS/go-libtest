# - Ruby - #

require 'json'
require_relative '../../libs/sssa-ruby/lib/sssa.rb'

input = File.read('./tmp/Ruby_Combine_Input.json')
args = JSON.parse(input)
o = File.new('./tmp/Ruby_Combine_Output.json', 'w')
o.write SSSA::combine(args).to_json
o.close
