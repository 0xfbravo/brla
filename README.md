
# brla

This repository is an unofficial `golang` wrapper for the [BRLA](https://brla.digital/) APIs.

- Official API documentation for standard users can be found [here](https://brla-account-api.readme.io/reference/welcome).
- For users with superuser roles, the documentation is available [here](https://brla-superuser-api.readme.io/reference/welcome).

This library is designed with **Clean Architecture principles** to enhance maintainability, testability, and scalability.

Please refer to the LICENSE file for licensing details.

---

## 📚 Methods Available

### 👨‍💻 Normal API

| Endpoint             | Description           | Status  |
|----------------------|-----------------------|---------|
| 🔜 Coming soon!      | Methods are not yet implemented. | ❌ |

### 👮 Superuser API

| Endpoint                       | Description                                                                                                                          | Status  |
|--------------------------------|--------------------------------------------------------------------------------------------------------------------------------------|---------|
| 🔑 Login                       | Authenticate using email and password. Returns a JWT token.                                                                          | ✔️      |
| Password Reset                 | Request a password reset.                                                                                                            | ❌      |
| Finish Password Reset          | Finish the password reset process.                                                                                                   | ❌      |
| Change account password        | Change the account password.                                                                                                         | ❌      |
| Logout                         | Invalidate the current JWT token.                                                                                                    | ❌      |
| 💸 Buy Static Pix              | Creates a ticket for minting BRLA. Duration can be configured, and extra fees can be paid using an BRLA Account balance if available | ✔️      |
| 💸 Sell                       | Create sell order with input information | ✔️      |
| Buy operations history         | Retrieves paginated BRLA buy transactions for user                                                                                   | ❌      |
| Pix to USD history             | Fetch Pix to USD conversion history                                                                                                  | ❌      |
| Pix to USD                     | Creates a ticket for buying a dollar stablecoin. Requires a valid token obtained from the fast-quote endpoint                        | ✔️      |
| 💰 Get Balance Of              | Retrieve blockchain balance                                                                                                          | ✔️      |
| 📜 Get Contract Addresses      | Fetch contract addresses                                                                                                             | ✔️      |
| 🛡️ Get KYC Status             | Check Know Your Customer status                                                                                                      | ✔️      |
| 🔓 Get Public Key              | Retrieve the public key                                                                                                              | ✔️      |
| 📉 Get Used Limit              | View the used transaction limit                                                                                                      | ✔️      |
| 🧑‍⚖️ KYC Level One (Natural) | Basic KYC for natural persons                                                                                                        | ✔️      |
| 🏢 KYC Level One (Legal)       | Basic KYC for legal entities                                                                                                         | ✔️      |
| 🧑‍⚖️ KYC Level Two (Natural) | Advanced KYC for natural persons                                                                                                     | ✔️      |
| 🏢 KYC Level Two (Legal)       | Advanced KYC for legal entities                                                                                                      | ✔️      |
| 🔜 Coming soon!                | More methods are not yet implemented.                                                                                                | ❌ |

---

## 🤝 Contributing

Contributions and suggestions are welcome!  
1. Fork this repository.  
2. Open a pull request with your changes.  
3. I’ll review it as soon as possible.  

---

## 🛠️ Linting

**TODO**: Add linting guidelines.

---

## 🙌 Special Thanks

Special thanks to [ThalesSathler](https://github.com/ThalesSathler) for inspiring this project and providing best practice examples.  

And, of course, to the [BRLA](https://brla.digital/) team, which is always available to help.  
