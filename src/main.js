const fs = require('fs')
const peg = require('pegjs')

var parser = peg.generate(fs.readFileSync('./yum.pegjs'))

parser.parse


