```bash
simple-game $ less setup.sh
simple-game $ ls -al setup.sh
-rw-r--r--  1 karuppiahn  staff  248 Dec  1 20:50 setup.sh
simple-game $ bash setup.sh
setup.sh: line 18: pnpm: command not found
simple-game $ vi setup.sh
simple-game $ bash setup.sh
Need to install the following packages:
  pnpm
Ok to proceed? (y) y
Lockfile is up-to-date, resolution step is skipped
Packages: +1543
++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
Packages are hard linked from the content-addressable store to the virtual store.
  Content-addressable store is at: /Users/karuppiahn/.pnpm-store/v3
  Virtual store is at:             node_modules/.pnpm
Progress: resolved 1543, reused 0, downloaded 1543, added 1543, done
node_modules/.pnpm/core-js-pure@3.15.2/node_modules/core-js-pure: Running postinstall script, done in 71ms
node_modules/.pnpm/core-js@2.6.12/node_modules/core-js: Running postinstall script, done in 76ms
node_modules/.pnpm/ejs@2.7.4/node_modules/ejs: Running postinstall script, done in 95ms
node_modules/.pnpm/fsevents@1.2.13/node_modules/fsevents: Running install script, done in 9.5s
node_modules/.pnpm/core-js@3.15.2/node_modules/core-js: Running postinstall script, done in 69ms
node_modules/.pnpm/node-sass@6.0.1/node_modules/node-sass: Running install script, done in 2.2s
node_modules/.pnpm/node-sass@6.0.1/node_modules/node-sass: Running postinstall script, done in 410ms

dependencies:
+ @laufire/resist 0.0.4
+ @laufire/utils 2.9.2
+ @testing-library/jest-dom 5.14.1
+ @testing-library/react 11.2.7
+ @testing-library/user-event 12.8.3
+ node-sass 6.0.1
+ react 17.0.2
+ react-dom 17.0.2
+ react-scripts 4.0.3
+ web-vitals 1.1.2

devDependencies:
+ eslint 7.29.0
+ eslint-plugin-react 7.24.0
npm notice
npm notice New major version of npm available! 7.21.0 -> 8.1.4
npm notice Changelog: https://github.com/npm/cli/releases/tag/v8.1.4
npm notice Run npm install -g npm@8.1.4 to update!
npm notice
simple-game $ npm i -g npm

removed 46 packages, changed 44 packages, and audited 223 packages in 4s

10 packages are looking for funding
  run `npm fund` for details

3 moderate severity vulnerabilities

To address all issues, run:
  npm audit fix

Run `npm audit` for details.
simple-game $ npm audit
# npm audit report

ansi-html  *
Severity: high
Uncontrolled Resource Consumption in ansi-html - https://github.com/advisories/GHSA-whgm-jr23-g3j9
fix available via `npm audit fix --force`
Will install react-scripts@3.4.4, which is a breaking change
node_modules/ansi-html
  @pmmmwh/react-refresh-webpack-plugin  <=0.5.0-rc.6
  Depends on vulnerable versions of ansi-html
  node_modules/react-scripts/node_modules/@pmmmwh/react-refresh-webpack-plugin
    react-scripts  >=0.10.0-alpha.328cb32e
    Depends on vulnerable versions of @pmmmwh/react-refresh-webpack-plugin
    Depends on vulnerable versions of @svgr/webpack
    Depends on vulnerable versions of optimize-css-assets-webpack-plugin
    Depends on vulnerable versions of react-dev-utils
    Depends on vulnerable versions of webpack
    Depends on vulnerable versions of webpack-dev-server
    node_modules/react-scripts
  webpack-dev-server  2.0.0-beta - 4.1.0
  Depends on vulnerable versions of ansi-html
  Depends on vulnerable versions of chokidar
  Depends on vulnerable versions of yargs
  node_modules/webpack-dev-server

ansi-regex  >2.1.1 <5.0.1
Severity: moderate
 Inefficient Regular Expression Complexity in chalk/ansi-regex - https://github.com/advisories/GHSA-93q8-gq69-wqmw
fix available via `npm audit fix --force`
Will install react-scripts@3.4.4, which is a breaking change
node_modules/ansi-regex
node_modules/sass-graph/node_modules/ansi-regex
node_modules/webpack-dev-server/node_modules/cliui/node_modules/ansi-regex
node_modules/webpack-dev-server/node_modules/string-width/node_modules/ansi-regex
node_modules/webpack-dev-server/node_modules/wrap-ansi/node_modules/ansi-regex
  strip-ansi  4.0.0 - 5.2.0
  Depends on vulnerable versions of ansi-regex
  node_modules/sass-graph/node_modules/strip-ansi
  node_modules/webpack-dev-server/node_modules/cliui/node_modules/strip-ansi
  node_modules/webpack-dev-server/node_modules/string-width/node_modules/strip-ansi
  node_modules/webpack-dev-server/node_modules/wrap-ansi/node_modules/strip-ansi
    cliui  4.0.0 - 5.0.0
    Depends on vulnerable versions of strip-ansi
    Depends on vulnerable versions of wrap-ansi
    node_modules/sass-graph/node_modules/cliui
    node_modules/webpack-dev-server/node_modules/cliui
      yargs  10.1.0 - 15.0.0
      Depends on vulnerable versions of cliui
      Depends on vulnerable versions of string-width
      node_modules/sass-graph/node_modules/yargs
      node_modules/webpack-dev-server/node_modules/yargs
        sass-graph  2.2.5 || >=3.0.5
        Depends on vulnerable versions of yargs
        node_modules/sass-graph
          node-sass  >=4.14.1
          Depends on vulnerable versions of sass-graph
          node_modules/node-sass
        webpack-dev-server  2.0.0-beta - 4.1.0
        Depends on vulnerable versions of ansi-html
        Depends on vulnerable versions of chokidar
        Depends on vulnerable versions of yargs
        node_modules/webpack-dev-server
          react-scripts  >=0.10.0-alpha.328cb32e
          Depends on vulnerable versions of @pmmmwh/react-refresh-webpack-plugin
          Depends on vulnerable versions of @svgr/webpack
          Depends on vulnerable versions of optimize-css-assets-webpack-plugin
          Depends on vulnerable versions of react-dev-utils
          Depends on vulnerable versions of webpack
          Depends on vulnerable versions of webpack-dev-server
          node_modules/react-scripts
    string-width  2.1.0 - 4.1.0
    Depends on vulnerable versions of strip-ansi
    node_modules/sass-graph/node_modules/string-width
    node_modules/webpack-dev-server/node_modules/string-width
      wrap-ansi  3.0.0 - 6.1.0
      Depends on vulnerable versions of string-width
      Depends on vulnerable versions of strip-ansi
      node_modules/sass-graph/node_modules/wrap-ansi
      node_modules/webpack-dev-server/node_modules/wrap-ansi

browserslist  4.0.0 - 4.16.4
Severity: moderate
Regular Expression Denial of Service in browserslist - https://github.com/advisories/GHSA-w8qv-6jwh-64r5
fix available via `npm audit fix --force`
Will install react-scripts@3.4.4, which is a breaking change
node_modules/react-dev-utils/node_modules/browserslist
  react-dev-utils  6.0.0-next.03604a46 - 12.0.0-next.37
  Depends on vulnerable versions of browserslist
  Depends on vulnerable versions of immer
  node_modules/react-dev-utils
    react-scripts  >=0.10.0-alpha.328cb32e
    Depends on vulnerable versions of @pmmmwh/react-refresh-webpack-plugin
    Depends on vulnerable versions of @svgr/webpack
    Depends on vulnerable versions of optimize-css-assets-webpack-plugin
    Depends on vulnerable versions of react-dev-utils
    Depends on vulnerable versions of webpack
    Depends on vulnerable versions of webpack-dev-server
    node_modules/react-scripts

glob-parent  <5.1.2
Severity: high
Regular expression denial of service - https://github.com/advisories/GHSA-ww39-953v-wcq6
fix available via `npm audit fix --force`
Will install react-scripts@3.4.4, which is a breaking change
node_modules/watchpack-chokidar2/node_modules/glob-parent
node_modules/webpack-dev-server/node_modules/glob-parent
  chokidar  1.0.0-rc1 - 2.1.8
  Depends on vulnerable versions of glob-parent
  node_modules/watchpack-chokidar2/node_modules/chokidar
  node_modules/webpack-dev-server/node_modules/chokidar
    watchpack-chokidar2  *
    Depends on vulnerable versions of chokidar
    node_modules/watchpack-chokidar2
      watchpack  1.7.2 - 1.7.5
      Depends on vulnerable versions of watchpack-chokidar2
      node_modules/watchpack
        webpack  4.44.0 - 4.46.0
        Depends on vulnerable versions of watchpack
        node_modules/webpack
          react-scripts  >=0.10.0-alpha.328cb32e
          Depends on vulnerable versions of @pmmmwh/react-refresh-webpack-plugin
          Depends on vulnerable versions of @svgr/webpack
          Depends on vulnerable versions of optimize-css-assets-webpack-plugin
          Depends on vulnerable versions of react-dev-utils
          Depends on vulnerable versions of webpack
          Depends on vulnerable versions of webpack-dev-server
          node_modules/react-scripts
    webpack-dev-server  2.0.0-beta - 4.1.0
    Depends on vulnerable versions of ansi-html
    Depends on vulnerable versions of chokidar
    Depends on vulnerable versions of yargs
    node_modules/webpack-dev-server

immer  <9.0.6
Severity: critical
Prototype Pollution in immer - https://github.com/advisories/GHSA-33f9-j839-rf8h
fix available via `npm audit fix --force`
Will install react-scripts@3.4.4, which is a breaking change
node_modules/immer
  react-dev-utils  6.0.0-next.03604a46 - 12.0.0-next.37
  Depends on vulnerable versions of browserslist
  Depends on vulnerable versions of immer
  node_modules/react-dev-utils
    react-scripts  >=0.10.0-alpha.328cb32e
    Depends on vulnerable versions of @pmmmwh/react-refresh-webpack-plugin
    Depends on vulnerable versions of @svgr/webpack
    Depends on vulnerable versions of optimize-css-assets-webpack-plugin
    Depends on vulnerable versions of react-dev-utils
    Depends on vulnerable versions of webpack
    Depends on vulnerable versions of webpack-dev-server
    node_modules/react-scripts

json-schema  <0.4.0
Severity: moderate
json-schema is vulnerable to Prototype Pollution - https://github.com/advisories/GHSA-896r-f27r-55mw
fix available via `npm audit fix`
node_modules/json-schema
  jsprim  0.3.0 - 1.4.1 || 2.0.0 - 2.0.1
  Depends on vulnerable versions of json-schema
  node_modules/jsprim

nth-check  <2.0.1
Severity: moderate
Inefficient Regular Expression Complexity in nth-check - https://github.com/advisories/GHSA-rp65-9cf3-cjxr
fix available via `npm audit fix --force`
Will install react-scripts@3.4.4, which is a breaking change
node_modules/nth-check
node_modules/svgo/node_modules/nth-check
  css-select  <=3.1.0
  Depends on vulnerable versions of nth-check
  node_modules/svgo/node_modules/css-select
    svgo  1.0.0 - 1.3.2
    Depends on vulnerable versions of css-select
    node_modules/svgo
      @svgr/plugin-svgo  <=5.5.0
      Depends on vulnerable versions of svgo
      node_modules/@svgr/plugin-svgo
        @svgr/webpack  4.0.0 - 5.5.0
        Depends on vulnerable versions of @svgr/plugin-svgo
        node_modules/@svgr/webpack
          react-scripts  >=0.10.0-alpha.328cb32e
          Depends on vulnerable versions of @pmmmwh/react-refresh-webpack-plugin
          Depends on vulnerable versions of @svgr/webpack
          Depends on vulnerable versions of optimize-css-assets-webpack-plugin
          Depends on vulnerable versions of react-dev-utils
          Depends on vulnerable versions of webpack
          Depends on vulnerable versions of webpack-dev-server
          node_modules/react-scripts
      postcss-svgo  4.0.0-nightly.2020.1.9 - 5.0.0-rc.2
      Depends on vulnerable versions of svgo
      node_modules/postcss-svgo
        cssnano-preset-default  <=4.0.8
        Depends on vulnerable versions of postcss-svgo
        node_modules/cssnano-preset-default
          cssnano  4.0.0-nightly.2020.1.9 - 4.1.11
          Depends on vulnerable versions of cssnano-preset-default
          node_modules/cssnano
            optimize-css-assets-webpack-plugin  3.2.1 || 5.0.0 - 5.0.8
            Depends on vulnerable versions of cssnano
            node_modules/optimize-css-assets-webpack-plugin

31 vulnerabilities (20 moderate, 9 high, 2 critical)

To address issues that do not require attention, run:
  npm audit fix

To address all issues (including breaking changes), run:
  npm audit fix --force
simple-game $
simple-game $ cat setup.sh
#!/bin/bash

set -e

# Init
cd $(dirname $0)

# Tasks
setupEnvironment() {
	hooksPath="./.githooks"

	git config core.hooksPath "$hooksPath"
	chmod 775 "$hooksPath"/*
}

installPackages() {

	npx pnpm install
}

# Main
setupEnvironment
installPackages
simple-game $
```

```bash
 PASS  src/services/helperService.test.js
 PASS  src/services/playerManger.test.js
 PASS  src/services/targetManager/index.test.js
 PASS  src/services/gameService.test.js
 PASS  src/core/actions.test.js
 PASS  src/services/positionService.test.js
 PASS  src/components/game.test.js
 PASS  src/components/gameScreen.test.js
 PASS  src/components/bullet.test.js
 PASS  src/components/target.test.js
 PASS  src/components/gameOverScreen.test.js
 PASS  src/components/healthBar.test.js
 PASS  src/components/score.test.js
 PASS  src/components/flight.test.js
 PASS  src/components/restart.test.js
 PASS  src/components/cloud.test.js
 PASS  src/components/container.test.js
 PASS  src/services/ticker/index.test.js
 PASS  src/services/ticker/masterLoop.test.js
 PASS  src/App.test.js

Test Suites: 20 passed, 20 total
Tests:       66 passed, 66 total
Snapshots:   0 total
Time:        6.794 s
Ran all test suites.

Watch Usage
 › Press f to run only failed tests.
 › Press o to only run tests related to changed files.
 › Press q to quit watch mode.
 › Press p to filter by a filename regex pattern.
 › Press t to filter by a test name regex pattern.
 › Press Enter to trigger a test run.

simple-game $ less package.json
simple-game $ less package.json
simple-game $ less package.json
simple-game $ npx pnpm test-ci

> react-starter@0.1.0 test-ci /Users/karuppiahn/projects/github.com/NirupanJoe/simple-game
> npm run test-dev -- --coverage


> react-starter@0.1.0 test-dev
> react-scripts test --watchAll=false "--coverage"

 PASS  src/services/helperService.test.js (6.899 s)
 PASS  src/services/gameService.test.js (6.99 s)
 PASS  src/services/targetManager/index.test.js (6.992 s)
 PASS  src/services/positionService.test.js (7.005 s)
 PASS  src/services/playerManger.test.js (7.105 s)
 PASS  src/core/actions.test.js (7.178 s)
 PASS  src/components/score.test.js (7.208 s)
 PASS  src/components/healthBar.test.js (7.261 s)
 PASS  src/components/target.test.js (7.263 s)
 PASS  src/components/bullet.test.js (7.257 s)
 PASS  src/components/gameOverScreen.test.js (7.34 s)
 PASS  src/components/game.test.js (7.405 s)
 PASS  src/components/restart.test.js (7.313 s)
 PASS  src/components/flight.test.js (7.344 s)
 PASS  src/components/gameScreen.test.js (7.463 s)
 PASS  src/components/container.test.js
 PASS  src/services/ticker/index.test.js
 PASS  src/services/ticker/masterLoop.test.js
 PASS  src/components/cloud.test.js
 PASS  src/App.test.js
----------------------------|---------|----------|---------|---------|-------------------
File                        | % Stmts | % Branch | % Funcs | % Lines | Uncovered Line #s
----------------------------|---------|----------|---------|---------|-------------------
All files                   |     100 |      100 |     100 |     100 |
 src                        |     100 |      100 |     100 |     100 |
  App.js                    |     100 |      100 |     100 |     100 |
 src/components             |     100 |      100 |     100 |     100 |
  bullet.js                 |     100 |      100 |     100 |     100 |
  cloud.js                  |     100 |      100 |     100 |     100 |
  container.js              |     100 |      100 |     100 |     100 |
  flight.js                 |     100 |      100 |     100 |     100 |
  game.js                   |     100 |      100 |     100 |     100 |
  gameOverScreen.js         |     100 |      100 |     100 |     100 |
  gameScreen.js             |     100 |      100 |     100 |     100 |
  healthBar.js              |     100 |      100 |     100 |     100 |
  restart.js                |     100 |      100 |     100 |     100 |
  score.js                  |     100 |      100 |     100 |     100 |
  target.js                 |     100 |      100 |     100 |     100 |
 src/core                   |     100 |      100 |     100 |     100 |
  actions.js                |     100 |      100 |     100 |     100 |
  config.js                 |     100 |      100 |     100 |     100 |
  context.js                |     100 |      100 |     100 |     100 |
  seed.js                   |     100 |      100 |     100 |     100 |
 src/services               |     100 |      100 |     100 |     100 |
  gameService.js            |     100 |      100 |     100 |     100 |
  helperService.js          |     100 |      100 |     100 |     100 |
  playerManger.js           |     100 |      100 |     100 |     100 |
  positionService.js        |     100 |      100 |     100 |     100 |
 src/services/targetManager |     100 |      100 |     100 |     100 |
  index.js                  |     100 |      100 |     100 |     100 |
 src/services/ticker        |     100 |      100 |     100 |     100 |
  index.js                  |     100 |      100 |     100 |     100 |
  masterLoop.js             |     100 |      100 |     100 |     100 |
----------------------------|---------|----------|---------|---------|-------------------

Test Suites: 20 passed, 20 total
Tests:       66 passed, 66 total
Snapshots:   0 total
Time:        11.309 s
Ran all test suites.
simple-game $
```
