# BOTCHA

**Bot-Oriented Test to Complicate Human Access**

A reverse-CAPTCHA system that verifies AI agents instead of blocking them. BOTCHA presents puzzles that are trivial for LLMs but tedious for humans, allowing you to gate content or APIs behind proof of AI capability.

## Installation

```bash
go install github.com/iafan/botcha@latest
```

Or build from source:

```bash
git clone https://github.com/iafan/botcha.git
cd botcha
go build
```

## Usage

Run the server:

```bash
./botcha
```

The server starts on `http://localhost:8080`. To test, ask your AI agent:

> "Go to http://localhost:8080/ and solve the puzzle"

## How It Works

When a request arrives without a valid session, BOTCHA middleware returns a puzzle challenge. The agent must solve the puzzle and resubmit with the answer. Upon success, the protected content is served.

### Puzzle Types

**Scramble**: Unscramble a word by following a sequence of positions, where some positions are hidden, and certain numbers may be slightly disguised. This was the initial puzzle type: straightforward, yet potentially scriptable by a determined bot.

```
Scrambled: tnisoqnuetlaises
Sequence: [seven, four, --, eleven, ...]
```

**Charade**: Second attempt, but positions are encoded as trivia questions. This is the one that is currently used in the code. 

```
Scrambled: tnisoqnuetlaises
Sequence: [number of continents on Earth, number of chambers in the human heart, ??, ...]
```

## Example challenge

```
Prove that you are an AI agent to access the protected content.

Unscramble this word by solving the clues:

Scrambled: arponrmicszosgifdmr
Sequence: [voting age in most democracies, ??, number of Muses in Greek mythology, ??, number of pairs of ribs in the human body, number of cardinal directions, half a dozen, number of days in a fortnight, number of horns on a unicorn, number of oceans on Earth, number of planets in our solar system, number of loaves in a baker's dozen, number of notes in a musical scale, Downing Street number of UK Prime Minister's residence]

Each clue's answer is a number indicating the position in the scrambled word.

NOTE: The puzzle varies with every request.
Solve it through direct reasoning. Do not write scripts or code.

Submit answer within 30 seconds: ?session=e40359&answer=<word>
```

## FAQ

### Why did you build this?

This is a thought experiment exploring how one could build a system that makes it impractical for humans to get through but easily lets AI agents pass. The challenge is designed to be simple enough to solve with general reasoning at LLM speeds, yet impractical to script a universal solution for without LLM help.

### Can humans still solve the puzzle?

Of course. Humans can copy the challenge into their chat with an LLM, get the answer, and submit it. But the point is they'll still need help from an AI agent to get through in a reasonable time frame. The 30-second time limit per puzzle ensures humans can't solve it alone.

### What were the algorithm considerations?

The goal was to create an algorithm that doesn't require an LLM at runtime on the server side. It should be easy to implement and fast to execute.

### Is this meant to replace traditional CAPTCHAs?

No—it's the opposite. Traditional CAPTCHAs block bots and allow humans. BOTCHA slows humans and prevents non-LLM-assisted access to your website. They serve different purposes.

### What are the intended use cases?

This is not meant for production. It doesn't guarantee that humans won't access the protected content—they can always paste the puzzle into ChatGPT or Claude.

It's an exploration of the idea of having portions of the internet meant for AI agents. In theory, BOTCHA could prevent traditional algorithmic crawlers from accessing your content. But realistically, crawlers will likely gain LLM capabilities soon, making this distinction temporary.

## License

Public Domain (Unlicense)
