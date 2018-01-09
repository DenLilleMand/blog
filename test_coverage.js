var fs = require('fs');
const badge = require('gh-badges')

const format = {
  text: ['build', 'passed'],
  colorscheme: 'green',
  template: 'flat',
}

badge.loadFont('/usr/share/fonts/TTF/Hack-Regular.ttf', err => {
  badge(format, (svg, err) => {
    // svg is a string containing your badge
    fs.writeFile('testcoverage.svg', svg, (err) => {
      if (err) throw err;
      console.log('The file has been saved!');
    });
  })})


