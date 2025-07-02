# Changelog

## 1.37.0 (2025-07-02)

Full Changelog: [v1.34.1...v1.37.0](https://github.com/dodopayments/dodopayments-go/compare/v1.34.1...v1.37.0)

### Features

* **api:** updated openapi spec for v1.37.0 ([d30877b](https://github.com/dodopayments/dodopayments-go/commit/d30877b15a9b2afbebec951cdf6a3d180dd92de8))


### Bug Fixes

* don't try to deserialize as json when ResponseBodyInto is []byte ([93f4181](https://github.com/dodopayments/dodopayments-go/commit/93f41811368755de1b155f350e9b318ff2e1455f))
* **pagination:** check if page data is empty in GetNextPage ([75db278](https://github.com/dodopayments/dodopayments-go/commit/75db2787fcf1a25fafb2179c6c02180bf2e2723e))


### Chores

* **ci:** only run for pushes and fork pull requests ([523798a](https://github.com/dodopayments/dodopayments-go/commit/523798a2af676b41e6abf1e63596e6b88d4cca8b))

## 1.34.1 (2025-06-21)

Full Changelog: [v1.34.0...v1.34.1](https://github.com/dodopayments/dodopayments-go/compare/v1.34.0...v1.34.1)

## 1.34.0 (2025-06-18)

Full Changelog: [v1.32.0...v1.34.0](https://github.com/dodopayments/dodopayments-go/compare/v1.32.0...v1.34.0)

### Features

* **api:** updated to version 1.34.0 ([5d11342](https://github.com/dodopayments/dodopayments-go/commit/5d11342e2f69e292141166f5ce8ca0ac592c2826))
* **client:** add debug log helper ([3dc8c47](https://github.com/dodopayments/dodopayments-go/commit/3dc8c47637fb9f0940375b4764fbffdd4557a087))


### Chores

* **ci:** enable for pull requests ([99a67e6](https://github.com/dodopayments/dodopayments-go/commit/99a67e6a049b9bade204c9043d0bf03430b8835f))

## 1.32.0 (2025-06-09)

Full Changelog: [v1.30.2...v1.32.0](https://github.com/dodopayments/dodopayments-go/compare/v1.30.2...v1.32.0)

### Features

* **api:** updated openapi spec to v1.32.0 ([b8a42ad](https://github.com/dodopayments/dodopayments-go/commit/b8a42add4a7b4762f7f6719ae93a34c1bbb240ea))

## 1.30.2 (2025-06-04)

Full Changelog: [v1.30.0...v1.30.2](https://github.com/dodopayments/dodopayments-go/compare/v1.30.0...v1.30.2)

### Features

* **api:** fixed openapi spec ([c16ab44](https://github.com/dodopayments/dodopayments-go/commit/c16ab442b86129bc70f65ce331a0b87a261a5bcb))

## 1.30.0 (2025-06-02)

Full Changelog: [v1.27.0...v1.30.0](https://github.com/dodopayments/dodopayments-go/compare/v1.27.0...v1.30.0)

### Features

* **api:** manual updates ([cd3ac15](https://github.com/dodopayments/dodopayments-go/commit/cd3ac15f6f4fd293580224f076c989515402bfbe))


### Chores

* make go mod tidy continue on error ([377ec0c](https://github.com/dodopayments/dodopayments-go/commit/377ec0cff8b734bd626a36f7c37e9b6ede76978f))

## 1.27.0 (2025-05-26)

Full Changelog: [v1.25.0...v1.27.0](https://github.com/dodopayments/dodopayments-go/compare/v1.25.0...v1.27.0)

### Features

* **api:** added brands api in our sdk ([b35aa67](https://github.com/dodopayments/dodopayments-go/commit/b35aa6765326ebfd27c881e59f2d7e77c79fbccb))
* **api:** updated openapi spec to 1.27.0 ([62964db](https://github.com/dodopayments/dodopayments-go/commit/62964db3918c3f764ea9c9a1b176c1f715ef8bbd))


### Chores

* **docs:** grammar improvements ([e9a332c](https://github.com/dodopayments/dodopayments-go/commit/e9a332c0e20d69d0c3b93330188bb9a4ab68bc2e))
* improve devcontainer setup ([112f5c2](https://github.com/dodopayments/dodopayments-go/commit/112f5c2bcb1913e34457968f2028004583399b41))

## 1.25.0 (2025-05-17)

Full Changelog: [v1.22.0...v1.25.0](https://github.com/dodopayments/dodopayments-go/compare/v1.22.0...v1.25.0)

### Features

* **api:** updated openapi spec ([eefc5f8](https://github.com/dodopayments/dodopayments-go/commit/eefc5f8f4427820a399177f55554d3e787ae635b))
* **client:** add support for endpoint-specific base URLs in python ([b27eaa8](https://github.com/dodopayments/dodopayments-go/commit/b27eaa81cec506e81092b6e70e5510759ea59f29))

## 1.22.0 (2025-05-09)

Full Changelog: [v1.20.0...v1.22.0](https://github.com/dodopayments/dodopayments-go/compare/v1.20.0...v1.22.0)

### Features

* **api:** fixed api key schema to bearer ([8f68c28](https://github.com/dodopayments/dodopayments-go/commit/8f68c288629378a35ee4cc37ab1a001ed47ca3bb))
* **api:** manual updates ([49fc0bc](https://github.com/dodopayments/dodopayments-go/commit/49fc0bcbd8baea3bf23e586e18b48c524e4eda1e))
* **api:** updated openapi spec ([7a5f7e9](https://github.com/dodopayments/dodopayments-go/commit/7a5f7e96e1fc121b811992b4a72787c1bc6237e0))


### Bug Fixes

* **client:** clean up reader resources ([18f1eb1](https://github.com/dodopayments/dodopayments-go/commit/18f1eb1cc94554185add08b6512953d0f3b71ced))
* **client:** correctly update body in WithJSONSet ([0791e69](https://github.com/dodopayments/dodopayments-go/commit/0791e69f6a3e4cfde4fa5e01319ec284bb405cdd))


### Chores

* **internal:** codegen related update ([7631a6e](https://github.com/dodopayments/dodopayments-go/commit/7631a6e94f8f4f91318215afbfacbb3e6c6740e3))

## 1.20.0 (2025-05-01)

Full Changelog: [v1.19.0...v1.20.0](https://github.com/dodopayments/dodopayments-go/compare/v1.19.0...v1.20.0)

### Features

* **api:** added addons ([e62c54a](https://github.com/dodopayments/dodopayments-go/commit/e62c54a3049f4e530c8bf7d31399bbcd4358069c))
* **api:** updated readme example ([17f7321](https://github.com/dodopayments/dodopayments-go/commit/17f732157a5968a1935a04368231611267e2ccd8))
* **api:** updated readme example ([0885db4](https://github.com/dodopayments/dodopayments-go/commit/0885db4292cc78543f30a3211acab035b4c96f28))

## 1.19.0 (2025-04-30)

Full Changelog: [v1.18.3...v1.19.0](https://github.com/dodopayments/dodopayments-go/compare/v1.18.3...v1.19.0)

### Features

* **api:** manual updates ([b39feb5](https://github.com/dodopayments/dodopayments-go/commit/b39feb55d041614f075715455434d3a3ad58680c))


### Bug Fixes

* handle empty bodies in WithJSONSet ([c4a512f](https://github.com/dodopayments/dodopayments-go/commit/c4a512f5e226afea4b2fb511e248e200be09f6cc))

## 1.18.3 (2025-04-25)

Full Changelog: [v1.18.1...v1.18.3](https://github.com/dodopayments/dodopayments-go/compare/v1.18.1...v1.18.3)

### Features

* **api:** manual updates ([ed657d9](https://github.com/dodopayments/dodopayments-go/commit/ed657d96f645508274c822d09c5ff9b920a53963))

## 1.18.1 (2025-04-24)

Full Changelog: [v1.18.0...v1.18.1](https://github.com/dodopayments/dodopayments-go/compare/v1.18.0...v1.18.1)

### Chores

* **ci:** only use depot for staging repos ([675733f](https://github.com/dodopayments/dodopayments-go/commit/675733f2a217e9f2cb6c11e5c5b02f579e176f17))
* **internal:** codegen related update ([13d98e5](https://github.com/dodopayments/dodopayments-go/commit/13d98e58ae49b29df988eec666dcd932225afc48))

## 1.18.0 (2025-04-23)

Full Changelog: [v1.17.0...v1.18.0](https://github.com/dodopayments/dodopayments-go/compare/v1.17.0...v1.18.0)

### Features

* **api:** added change plan api ([680c698](https://github.com/dodopayments/dodopayments-go/commit/680c698c4e5a2bf90fdb74e953ac548667916ae2))
* **api:** manual updates ([ff481c9](https://github.com/dodopayments/dodopayments-go/commit/ff481c94c1bcfc6dbd48de598bcd84ce26f57149))


### Chores

* **ci:** add timeout thresholds for CI jobs ([65dace7](https://github.com/dodopayments/dodopayments-go/commit/65dace711d0670213088c9ac0aeaa4ba545eaa92))

## 1.17.0 (2025-04-22)

Full Changelog: [v1.16.1...v1.17.0](https://github.com/dodopayments/dodopayments-go/compare/v1.16.1...v1.17.0)

### Features

* **api:** manual updates ([c50a143](https://github.com/dodopayments/dodopayments-go/commit/c50a143039d8aaf12250313a9e114194b3194238))

## 1.16.1 (2025-04-18)

Full Changelog: [v1.14.2...v1.16.1](https://github.com/dodopayments/dodopayments-go/compare/v1.14.2...v1.16.1)

### Features

* **api:** manual updates ([a8c1288](https://github.com/dodopayments/dodopayments-go/commit/a8c12886836a34a7203eff023151beb20e8fb5fb))

## 1.14.2 (2025-04-17)

Full Changelog: [v1.14.1...v1.14.2](https://github.com/dodopayments/dodopayments-go/compare/v1.14.1...v1.14.2)

### Chores

* configure new SDK language ([9f96b13](https://github.com/dodopayments/dodopayments-go/commit/9f96b1376a5dec7bf065e9cfd3dfd4634ec9f6ac))

## 1.14.1 (2025-04-17)

Full Changelog: [v1.14.0...v1.14.1](https://github.com/dodopayments/dodopayments-go/compare/v1.14.0...v1.14.1)

### Features

* **client:** add support for reading base URL from environment variable ([6657397](https://github.com/dodopayments/dodopayments-go/commit/665739756fa0ded950f8626a0457f133d257afeb))


### Chores

* **docs:** document pre-request options ([c9e2855](https://github.com/dodopayments/dodopayments-go/commit/c9e285561788daeb2e473f6a9bafa4560b87a335))


### Documentation

* update documentation links to be more uniform ([deb7632](https://github.com/dodopayments/dodopayments-go/commit/deb763245f79c62e1cd0db30c6a2ce7f0d65d14a))

## 1.14.0 (2025-04-11)

Full Changelog: [v1.13.0...v1.14.0](https://github.com/dodopayments/dodopayments-go/compare/v1.13.0...v1.14.0)

### Features

* **api:** fixed license key pagination issues in openapi spec ([7a94e11](https://github.com/dodopayments/dodopayments-go/commit/7a94e11b04f3f0ada8518df06b9227216cde6e89))
* **api:** updated openapi spec ([79ea42b](https://github.com/dodopayments/dodopayments-go/commit/79ea42bdcc769eacf61e6013ee3cf213ba19a9d5))


### Chores

* **internal:** expand CI branch coverage ([0fa813a](https://github.com/dodopayments/dodopayments-go/commit/0fa813a6008d1c566d63416261f7ef7572eb302b))
* **internal:** reduce CI branch coverage ([70875a9](https://github.com/dodopayments/dodopayments-go/commit/70875a9c4054caa61eac8bbc6fe871c2e658a6db))

## 1.13.0 (2025-04-08)

Full Changelog: [v1.11.1...v1.13.0](https://github.com/dodopayments/dodopayments-go/compare/v1.11.1...v1.13.0)

### Features

* **api:** manual updates ([#103](https://github.com/dodopayments/dodopayments-go/issues/103)) ([0dc45f8](https://github.com/dodopayments/dodopayments-go/commit/0dc45f872e78b01f6349636b772f48cff3607bb1))
* **client:** support custom http clients ([#101](https://github.com/dodopayments/dodopayments-go/issues/101)) ([1eb6207](https://github.com/dodopayments/dodopayments-go/commit/1eb6207facbfa22c85d4a492a72cdf8cf4eac9d1))

## 1.11.1 (2025-04-05)

Full Changelog: [v1.11.0...v1.11.1](https://github.com/dodopayments/dodopayments-go/compare/v1.11.0...v1.11.1)

### Bug Fixes

* **client:** return error on bad custom url instead of panic ([#99](https://github.com/dodopayments/dodopayments-go/issues/99)) ([ed014bd](https://github.com/dodopayments/dodopayments-go/commit/ed014bdef82cb23fe4d1bb3269efe9732f883a0f))

## 1.11.0 (2025-03-28)

Full Changelog: [v1.10.4...v1.11.0](https://github.com/dodopayments/dodopayments-go/compare/v1.10.4...v1.11.0)

### Features

* **api:** manual updates ([#95](https://github.com/dodopayments/dodopayments-go/issues/95)) ([10fd546](https://github.com/dodopayments/dodopayments-go/commit/10fd546f1ad499ceab8fe56fb5dbcba01599a99f))

## 1.10.4 (2025-03-28)

Full Changelog: [v1.10.3...v1.10.4](https://github.com/dodopayments/dodopayments-go/compare/v1.10.3...v1.10.4)

### Bug Fixes

* **test:** return early after test failure ([#92](https://github.com/dodopayments/dodopayments-go/issues/92)) ([34fbb82](https://github.com/dodopayments/dodopayments-go/commit/34fbb82a253dc8929cdcd7ea55df5e27c9cd85b4))


### Chores

* add request options to client tests ([#90](https://github.com/dodopayments/dodopayments-go/issues/90)) ([c40f8e9](https://github.com/dodopayments/dodopayments-go/commit/c40f8e94da595050a56f3bfa74c432b69d53e672))
* fix typos ([#93](https://github.com/dodopayments/dodopayments-go/issues/93)) ([8f9e245](https://github.com/dodopayments/dodopayments-go/commit/8f9e245dc32500d6c25a11f423b8f7f5902c49c3))

## 1.10.3 (2025-03-25)

Full Changelog: [v1.10.1...v1.10.3](https://github.com/dodopayments/dodopayments-go/compare/v1.10.1...v1.10.3)

### Features

* **api:** manual updates ([#88](https://github.com/dodopayments/dodopayments-go/issues/88)) ([453a530](https://github.com/dodopayments/dodopayments-go/commit/453a53076a88d87f1322f261b801fb2a25676846))


### Chores

* **docs:** improve security documentation ([#86](https://github.com/dodopayments/dodopayments-go/issues/86)) ([29c6a85](https://github.com/dodopayments/dodopayments-go/commit/29c6a85f04866ba20bfdfecd0497387763aacc41))

## 1.10.1 (2025-03-21)

Full Changelog: [v1.7.0...v1.10.1](https://github.com/dodopayments/dodopayments-go/compare/v1.7.0...v1.10.1)

### Features

* **api:** updated openapispec to v1.10.1 ([#83](https://github.com/dodopayments/dodopayments-go/issues/83)) ([f396b68](https://github.com/dodopayments/dodopayments-go/commit/f396b6817031dfabb57865a3f8883a5456ea17d3))

## 1.7.0 (2025-03-14)

Full Changelog: [v1.6.3...v1.7.0](https://github.com/dodopayments/dodopayments-go/compare/v1.6.3...v1.7.0)

### Features

* **api:** fixed openapi spec issues ([#80](https://github.com/dodopayments/dodopayments-go/issues/80)) ([69727b9](https://github.com/dodopayments/dodopayments-go/commit/69727b9950f8953200070c6f9379c6e5007042c8))

## 1.6.3 (2025-03-14)

Full Changelog: [v1.5.1...v1.6.3](https://github.com/dodopayments/dodopayments-go/compare/v1.5.1...v1.6.3)

### Features

* **api:** openapi spec updated ([#77](https://github.com/dodopayments/dodopayments-go/issues/77)) ([45cc3ba](https://github.com/dodopayments/dodopayments-go/commit/45cc3baa9617699d0a5dec55eb2fa1ce48cbd676))
* **api:** updated stainless config ([#78](https://github.com/dodopayments/dodopayments-go/issues/78)) ([5b42536](https://github.com/dodopayments/dodopayments-go/commit/5b425369e01693925e5e7ae76625de401a2a80b3))
* **client:** improve default client options support ([#74](https://github.com/dodopayments/dodopayments-go/issues/74)) ([e965c60](https://github.com/dodopayments/dodopayments-go/commit/e965c6044741a43205028d2968311ce9eac8c72e))


### Chores

* **internal:** remove extra empty newlines ([#76](https://github.com/dodopayments/dodopayments-go/issues/76)) ([7597ba9](https://github.com/dodopayments/dodopayments-go/commit/7597ba94050f9b237b72a4580c22f4fe290c7185))

## 1.5.1 (2025-03-12)

Full Changelog: [v1.5.0...v1.5.1](https://github.com/dodopayments/dodopayments-go/compare/v1.5.0...v1.5.1)

### Features

* **client:** allow custom baseurls without trailing slash ([#70](https://github.com/dodopayments/dodopayments-go/issues/70)) ([26e932f](https://github.com/dodopayments/dodopayments-go/commit/26e932f83d6ded011e04e51dbcd5681c8cf367ce))


### Chores

* **internal:** codegen related update ([#72](https://github.com/dodopayments/dodopayments-go/issues/72)) ([ebd029c](https://github.com/dodopayments/dodopayments-go/commit/ebd029ce62659312c4ea66ee67f7124ff3cd2e7d))

## 1.5.0 (2025-03-07)

Full Changelog: [v1.0.0...v1.5.0](https://github.com/dodopayments/dodopayments-go/compare/v1.0.0...v1.5.0)

### Features

* **api:** manual updates ([#68](https://github.com/dodopayments/dodopayments-go/issues/68)) ([5966b33](https://github.com/dodopayments/dodopayments-go/commit/5966b33933679146bebefe17c0150993f7fe2f08))


### Chores

* use strconv.ParseInt instead of strconv.Atoi for pagination params ([#67](https://github.com/dodopayments/dodopayments-go/issues/67)) ([c69b09d](https://github.com/dodopayments/dodopayments-go/commit/c69b09d3022e892be271cb137c371a2ec0363845))


### Documentation

* update URLs from stainlessapi.com to stainless.com ([#65](https://github.com/dodopayments/dodopayments-go/issues/65)) ([52a0d12](https://github.com/dodopayments/dodopayments-go/commit/52a0d1226c64724a704c0bc828e847e7f300768c))

## 1.0.0 (2025-02-23)

Full Changelog: [v0.24.0...v1.0.0](https://github.com/dodopayments/dodopayments-go/compare/v0.24.0...v1.0.0)

### Features

* **api:** updated config and updated version to v1.0.0 ([#63](https://github.com/dodopayments/dodopayments-go/issues/63)) ([6aeee58](https://github.com/dodopayments/dodopayments-go/commit/6aeee580859e4bc3b4e7e20945b0479914fad2d7))


### Chores

* **internal:** codegen related update ([#61](https://github.com/dodopayments/dodopayments-go/issues/61)) ([6d0094f](https://github.com/dodopayments/dodopayments-go/commit/6d0094f97fbdf6fc8c7cf3a882608990005f84c5))

## 0.24.0 (2025-02-15)

Full Changelog: [v0.22.1...v0.24.0](https://github.com/dodopayments/dodopayments-go/compare/v0.22.1...v0.24.0)

### Features

* **api:** added discount apis ([#59](https://github.com/dodopayments/dodopayments-go/issues/59)) ([e05b567](https://github.com/dodopayments/dodopayments-go/commit/e05b56738825f46f54be61c11f8f24cd9b66d1c9))
* **api:** openapi spec changes ([#58](https://github.com/dodopayments/dodopayments-go/issues/58)) ([158f08f](https://github.com/dodopayments/dodopayments-go/commit/158f08f2eb9498764b0a2733fee4f6efad977f26))


### Bug Fixes

* **client:** don't truncate manually specified filenames ([#57](https://github.com/dodopayments/dodopayments-go/issues/57)) ([141d409](https://github.com/dodopayments/dodopayments-go/commit/141d4098b8ec71fbea19bbffae09f76fff2c7383))
* do not call path.Base on ContentType ([#55](https://github.com/dodopayments/dodopayments-go/issues/55)) ([5632b38](https://github.com/dodopayments/dodopayments-go/commit/5632b38d8c6da3829adb93ec340339a45dc2358b))

## 0.22.1 (2025-02-11)

Full Changelog: [v0.22.0...v0.22.1](https://github.com/dodopayments/dodopayments-go/compare/v0.22.0...v0.22.1)

### Features

* **api:** manual updates ([#53](https://github.com/dodopayments/dodopayments-go/issues/53)) ([9e33be2](https://github.com/dodopayments/dodopayments-go/commit/9e33be29da04c57bccd00a710b8eb4eeca07d0ca))


### Bug Fixes

* fix early cancel when RequestTimeout is provided for streaming requests ([#51](https://github.com/dodopayments/dodopayments-go/issues/51)) ([aba5c84](https://github.com/dodopayments/dodopayments-go/commit/aba5c84542dedea35df5742259ae4ef5d8959e1c))

## 0.22.0 (2025-02-06)

Full Changelog: [v0.20.2...v0.22.0](https://github.com/dodopayments/dodopayments-go/compare/v0.20.2...v0.22.0)

### Features

* **api:** updated API changes for v0.22.0 ([#49](https://github.com/dodopayments/dodopayments-go/issues/49)) ([4227b4e](https://github.com/dodopayments/dodopayments-go/commit/4227b4e9f4e9b52e23e771332cad77a29f9130b4))
* **client:** send `X-Stainless-Timeout` header ([#46](https://github.com/dodopayments/dodopayments-go/issues/46)) ([820c6d1](https://github.com/dodopayments/dodopayments-go/commit/820c6d165a7eecffda381ccfa706975627d151d8))


### Chores

* add UnionUnmarshaler for responses that are interfaces ([#48](https://github.com/dodopayments/dodopayments-go/issues/48)) ([c9c8181](https://github.com/dodopayments/dodopayments-go/commit/c9c81816a27fc6fcb8d899bb58471bd28544bf45))

## 0.20.2 (2025-02-01)

Full Changelog: [v0.20.1...v0.20.2](https://github.com/dodopayments/dodopayments-go/compare/v0.20.1...v0.20.2)

### Bug Fixes

* fix unicode encoding for json ([#42](https://github.com/dodopayments/dodopayments-go/issues/42)) ([315b144](https://github.com/dodopayments/dodopayments-go/commit/315b144ae3cfe253451c5dcb1ed0a291f5edcace))


### Documentation

* document raw responses ([#44](https://github.com/dodopayments/dodopayments-go/issues/44)) ([4ecb25e](https://github.com/dodopayments/dodopayments-go/commit/4ecb25eb69520af6ac6e77694aebd71bc3f8792a))

## 0.20.1 (2025-01-29)

Full Changelog: [v0.19.0...v0.20.1](https://github.com/dodopayments/dodopayments-go/compare/v0.19.0...v0.20.1)

### Features

* **api:** manual updates ([#40](https://github.com/dodopayments/dodopayments-go/issues/40)) ([63f43f0](https://github.com/dodopayments/dodopayments-go/commit/63f43f0d8f540ba057aaf887ed27b2f98ac6a6fe))


### Chores

* refactor client tests ([#38](https://github.com/dodopayments/dodopayments-go/issues/38)) ([2c48da1](https://github.com/dodopayments/dodopayments-go/commit/2c48da169b57653544caf7e0041dc5f38a4a386e))

## 0.19.0 (2025-01-23)

Full Changelog: [v0.18.0...v0.19.0](https://github.com/dodopayments/dodopayments-go/compare/v0.18.0...v0.19.0)

### Features

* **api:** added archive product api ([#31](https://github.com/dodopayments/dodopayments-go/issues/31)) ([c35157f](https://github.com/dodopayments/dodopayments-go/commit/c35157fa4b6fc5c4a1dc88b5600fb0cf0f5fadbc))
* **api:** manual updates ([#35](https://github.com/dodopayments/dodopayments-go/issues/35)) ([21343c9](https://github.com/dodopayments/dodopayments-go/commit/21343c9d032d598640b5a6ae765840669b9556a1))
* **api:** manual updates ([#36](https://github.com/dodopayments/dodopayments-go/issues/36)) ([768e7ad](https://github.com/dodopayments/dodopayments-go/commit/768e7adf65e57bebcd7444c7a88f362fdcc1ad4c))


### Bug Fixes

* fix apijson.Port for embedded structs ([#33](https://github.com/dodopayments/dodopayments-go/issues/33)) ([27f05fb](https://github.com/dodopayments/dodopayments-go/commit/27f05fb18f90fa5ca4a8569cac9f93c9077252cc))
* fix apijson.Port for embedded structs ([#34](https://github.com/dodopayments/dodopayments-go/issues/34)) ([bc00fd4](https://github.com/dodopayments/dodopayments-go/commit/bc00fd4d4689e7d92c945fbae5ea0b4062516e63))

## 0.18.0 (2025-01-20)

Full Changelog: [v0.17.0...v0.18.0](https://github.com/dodopayments/dodopayments-go/compare/v0.17.0...v0.18.0)

### Features

* **api:** updated openapi sepc ([#28](https://github.com/dodopayments/dodopayments-go/issues/28)) ([45929e2](https://github.com/dodopayments/dodopayments-go/commit/45929e2b4c0050d95253ba43243d912ed2265bda))

## 0.17.0 (2025-01-16)

Full Changelog: [v0.16.1...v0.17.0](https://github.com/dodopayments/dodopayments-go/compare/v0.16.1...v0.17.0)

### Features

* **api:** added filter apis ([#25](https://github.com/dodopayments/dodopayments-go/issues/25)) ([6a4478d](https://github.com/dodopayments/dodopayments-go/commit/6a4478d67c115aa5b918ad284c0629f7ae979a59))

## 0.16.1 (2025-01-11)

Full Changelog: [v0.15.1...v0.16.1](https://github.com/dodopayments/dodopayments-go/compare/v0.15.1...v0.16.1)

### Features

* **api:** updated openapi spec ([#23](https://github.com/dodopayments/dodopayments-go/issues/23)) ([0afaa8f](https://github.com/dodopayments/dodopayments-go/commit/0afaa8fe5c0301fd6f1edc492115fe8ebf6477dc))


### Chores

* **internal:** codegen related update ([#21](https://github.com/dodopayments/dodopayments-go/issues/21)) ([fc19638](https://github.com/dodopayments/dodopayments-go/commit/fc19638e88ba88f5fe4410bc68e6ac4b754bb657))

## 0.15.1 (2025-01-03)

Full Changelog: [v0.14.1...v0.15.1](https://github.com/dodopayments/dodopayments-go/compare/v0.14.1...v0.15.1)

### Features

* **api:** added invoice api and update openapi spec ([#19](https://github.com/dodopayments/dodopayments-go/issues/19)) ([d7ed6f9](https://github.com/dodopayments/dodopayments-go/commit/d7ed6f909937e21ba8b606b6c411502e18707d09))


### Chores

* **internal:** codegen related update ([#17](https://github.com/dodopayments/dodopayments-go/issues/17)) ([1dc57c4](https://github.com/dodopayments/dodopayments-go/commit/1dc57c46e58bb899a2b79aad047e7b4e9f857815))

## 0.14.1 (2024-12-29)

Full Changelog: [v0.14.0...v0.14.1](https://github.com/dodopayments/dodopayments-go/compare/v0.14.0...v0.14.1)

### Features

* **api:** manual updates ([#14](https://github.com/dodopayments/dodopayments-go/issues/14)) ([385fd53](https://github.com/dodopayments/dodopayments-go/commit/385fd53b0647839a6c347e3fd7899a2eb26414e2))

## 0.14.0 (2024-12-25)

Full Changelog: [v0.13.2...v0.14.0](https://github.com/dodopayments/dodopayments-go/compare/v0.13.2...v0.14.0)

### Features

* **api:** updated openapi spec for License Keys ([#11](https://github.com/dodopayments/dodopayments-go/issues/11)) ([de4f0cb](https://github.com/dodopayments/dodopayments-go/commit/de4f0cbe4b80b5909a00e6b4b01314a1b79a31d0))

## 0.13.2 (2024-12-21)

Full Changelog: [v0.12.0...v0.13.2](https://github.com/dodopayments/dodopayments-go/compare/v0.12.0...v0.13.2)

### Chores

* **internal:** codegen related update ([#8](https://github.com/dodopayments/dodopayments-go/issues/8)) ([a84732d](https://github.com/dodopayments/dodopayments-go/commit/a84732d9c90f3921c995cbe9f239ed1485f4ac69))

## 0.12.0 (2024-12-17)

Full Changelog: [v0.11.1...v0.12.0](https://github.com/dodopayments/dodopayments-go/compare/v0.11.1...v0.12.0)

### Features

* **api:** api update ([#4](https://github.com/dodopayments/dodopayments-go/issues/4)) ([d16f93a](https://github.com/dodopayments/dodopayments-go/commit/d16f93a17a91fdc1bb4549f347c1ffb11cc1a849))
* **api:** updated openapi methods ([#6](https://github.com/dodopayments/dodopayments-go/issues/6)) ([ad682de](https://github.com/dodopayments/dodopayments-go/commit/ad682de9a3cf90c61423b4c048d2aedfd4e41c92))

## 0.11.1 (2024-12-16)

Full Changelog: [v0.0.1-alpha.0...v0.11.1](https://github.com/dodopayments/dodopayments-go/compare/v0.0.1-alpha.0...v0.11.1)

### Features

* **api:** update via SDK Studio ([3030a4e](https://github.com/dodopayments/dodopayments-go/commit/3030a4eb16daa9093803c344ebde44971dd23dbf))
* **api:** update via SDK Studio ([d55f912](https://github.com/dodopayments/dodopayments-go/commit/d55f9127ee6a387b86ad98b5631504ec53e304dc))
* **api:** update via SDK Studio ([0ee6ea9](https://github.com/dodopayments/dodopayments-go/commit/0ee6ea90cd0bc866b76e2582c105056a06b09d6f))
* **api:** update via SDK Studio ([de430a4](https://github.com/dodopayments/dodopayments-go/commit/de430a447bdb48787f72e54cce9e6cc1b589fa44))
* **api:** update via SDK Studio ([77666d6](https://github.com/dodopayments/dodopayments-go/commit/77666d6e7051e6374b45bfca909ae4104ca84546))
* **api:** update via SDK Studio ([4206df9](https://github.com/dodopayments/dodopayments-go/commit/4206df9eac51b6ec1ccb1537c56db98f092cc899))
* **api:** update via SDK Studio ([26fc816](https://github.com/dodopayments/dodopayments-go/commit/26fc816430d73f0cc361a36d3954e7a782e23690))


### Chores

* configure new SDK language ([3234950](https://github.com/dodopayments/dodopayments-go/commit/3234950c15a41f39374cc21dd797ce5ca36b9465))
* go live ([#1](https://github.com/dodopayments/dodopayments-go/issues/1)) ([f57eaa9](https://github.com/dodopayments/dodopayments-go/commit/f57eaa9aca5cc85f3808e864a1c26966b4b7b609))
