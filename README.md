# Simple Bank

## Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
<!-- - [Contributing](../CONTRIBUTING.md) -->


 <!-- irm get.scoop.sh | iex -- Install scoop
 scoop install migrate  -- Install migrate -->
## About <a name = "about"></a>


It's practice project of simple banking system. where you can do the banking operation like Create Account, Transfer money account to account and make the entries of that transaction in entries.

## Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. 
<!-- See [deployment](#deployment) for notes on how to deploy the project on a live system. -->


### Prerequisites
These instructions are for Windows system. 
What things you need to install the software and how to install them.

You need to install several softwares for running the project in your local machine.

1. Docker
2. Postgres image   ```--version 12 alpine```
3. Scoop
4. Make
5. sqlc 
6. Migrate


### Installing

A step by step series of examples that tell you how to get a development env running.

Install the Docker Engine in your window system, to run the postgres image which provides you database for project.

Then pull and run the postgres image in your local. 

```
docker pull postgres:12-alpine

```
All the other commads regarding Docker is given into the make file.

Istall scoop to dowmload the Make and Migrate

```
irm get.scoop.sh | iex
```
Install Migrate & Make using scoop.

-- Install migrate
```
scoop install migrate       
```

-- Install make
```
scoop install make          
```

Rest are given in make file in project.

E.g.
```
make postgres
```

Check with make file in project for further update commands.

## Usage <a name = "usage"></a>

Start the docker engine and run the postgres image. Then start the banking operations with code. 
