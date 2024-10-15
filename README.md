<p align="center"><img src="https://imgur.com/Vp4LWt0.png" width=160 title="SIAM-VIT" alt="SIAM-VIT"></a>
</p>
<div align="center">
  <h3 align="center">MPL Backend</h3>

  <p align="center">
    <a href="https://github.com/orgs/SIAM-VIT/repositories?q=mpl"><strong>Explore other MPL repositories</strong></a>
    <br />
    <br />
    <a href="https://github.com/SIAM-VIT/MPL-treasurehunt-be/issues">Report Bug</a>
    ·
    <a href="https://github.com/SIAM-VIT/MPL-treasurehunt-be/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->

## Table of Contents

- [About The Project](#about-the-project)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Contributors](#contributors)
- [License](#license)

<!-- ABOUT THE PROJECT -->

## About The Project

This is the backend API for the Math Premier League (MPL), SIAM-VIT's Flagship Event in VIT's fest, GraVITas. This API is used to manage the check of the answer entered by the user in the clues. The API is built using Go Lang Fiber.

### Built With

This project is built using the following technologies and frameworks:

- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/docs/latest/)

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running, follow these simple steps.

### Prerequisites

Make sure you have the following installed:

- [Go](https://golang.org/doc/install)
- PostgreSQL 13 or higher
- Docker (optional but recommended)

### Installation

1. Clone the repo

   ```sh
   git clone https://github.com/SIAM-VIT/MPL-be.git

   ```

2. Change into the project directory

   ```sh
   cd MPL-be

   ```

3. Install dependencies

   ```sh
   go mod download

   ```

4. Set up PostgreSQL and configure .env with your database credentials:

   ```sh
    # Postgres Configuration
    DB_HOST=abc
    DB_PORT=abc
    DB_USER=abc
    DB_PASSWORD=abc
    DB_NAME=abc

   ```

5. Start the application
   ```sh
   go run cmd/main.go
   ```

## Contributors

<table>
	<tr align="center" style="font-weight:bold">
		<td>
		Akshat Gupta
		<p align="center">
			<img src = "https://avatars.githubusercontent.com/u/84951451?v=4" width="150" height="150" alt="Akshat Gupta">
		</p>
			<p align="center">
				<a href = "https://github.com/Oik17">
					<img src = "http://www.iconninja.com/files/241/825/211/round-collaboration-social-github-code-circle-network-icon.svg" width="36" height = "36" alt="GitHub"/>
				</a>
			</p>
		</td>
	</tr>
</table>

<!-- LICENSE -->

## License

Distributed under the MIT License. See LICENSE for more information.

<p align="center">
	Made with :heart: by SIAM-VIT
</p>
