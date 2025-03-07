# Changelog

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
