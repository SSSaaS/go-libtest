# - Ruby - #

require 'json'
require_relative '../../libs/sssa-ruby/lib/sssa.rb'

input = File.read('./tmp/Ruby_Create_Input.json')
args = JSON.parse(input)
o = File.new('./tmp/Ruby_Create_Output.json', 'w')
o.write SSSA::create(args[0].to_i, args[1].to_i, args[2]).to_json
o.close
