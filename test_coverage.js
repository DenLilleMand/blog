var fs = require('fs');
const badge = require('gh-badges')

fs.readFile('test_coverage.txt', (err, data) => {
  if (err) {
   throw err;   
  }
  //TODO consider the robostness of this
  var splitedTestCov = data.toString().split(':');
  //TODO maybe change the color scheme to red if below like 80%
  const format = {
    text: [splitedTestCov[0], splitedTestCov[1]],
    colorscheme: 'green',
    template: 'flat',
  }
  badge.loadFont('/usr/share/fonts/TTF/Hack-Regular.ttf', err => {
    badge(format, (svg, err) => {
    // svg is a string containing your badge
    fs.writeFile('testcoverage.svg', svg, (err) => {
      if (err) {
	throw err;  
      }
      console.log('The test_coverage.svg file has been saved. \n');
    });
  })})
});
