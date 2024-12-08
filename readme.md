# Catty Cuteness Voting

Welcome to **Catty Cuteness Voting**, a fun web application where users can cast their votes for the cutest cats! This project serves as both a fun platform for cat enthusiasts and a testing ground for the **Hawkwing** framework, which I developed to simplify web programming in Go.

Inspired by the webapp depicted in the movie **The Social Network**, this app allows users to compare two adorable feline contenders at a time, fostering a sense of rivalry and excitement. Each vote influences the dynamic rating system, which ranks our furry friends using an Elo-inspired algorithm. Just like in competitive chess, where players' skills are measured and adjusted based on their performance, _our cats’ ratings evolve with each interaction, reflecting their popularity among voters_.

<p align="center">
  <img src="assets/cattycuteness1.gif" alt="Voting Demo" width="600">
</p>
<p align="center">
  <em>CattyCuteness in action!</em>
</p>

## Features

- **Vote for Cats**: Users can vote for their favorite cats displayed on the homepage.
- **Dynamic Leaderboard**: The leaderboard updates in real-time to show each cat's rating based on user votes.
- **Two-at-a-Time Comparison**: Engage in friendly competition by voting between two randomly selected cats.
- **Responsive Design**: The application is designed to be fully responsive, providing a great user experience on both desktop and mobile devices.
- **Built with Hawkwing**: This application proudly utilizes the **Hawkwing** framework, a powerful tool I developed to streamline web programming in Go. Hawkwing simplifies the complexities of building web applications, offering intuitive routing, flexible middleware support, and robust templating features. It empowers developers to create dynamic and efficient web experiences with ease, making it a perfect fit for this cat voting app.

## Installation

To run the Catty Cuteness Voting application locally, follow these steps:

1. **Clone the Repository**:

```bash
git clone git@github.com:Aliqyan-21/CattyCuteness.git
cd CattyCuteness
```

2. Install Go:
   Make sure you have Go installed on your machine. You can download it from golang.org.

3. Run the Application:

```bash
go run main.go
```

## Usage

1. On the homepage, you will see two random cat images.
2. Click on the image of the cat you find cuter to cast your vote.
3. The leaderboard will update automatically to reflect the new ratings of each cat based on user votes.

## How It Works

The application employs an Elo-style rating system to dynamically update cat ratings based on user votes. Each time a vote is cast, the ratings of both cats are recalculated according to their expected scores and actual outcomes, creating an engaging experience reminiscent of competitive environments.

## Technologies Used

- Go (Golang)
- Hawkwing Framework
- HTML/CSS for frontend design

## Hawkwing Framework

The Hawkwing framework is designed to make web programming in Go easier and more efficient. It provides essential features for building web applications, such as routing, middleware support, and templating capabilities. You can explore more about it and contribute to its development [Here](https://github.com/aliqyan-21/hawkwing).
