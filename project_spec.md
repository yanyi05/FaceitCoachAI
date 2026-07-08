# FaceitCoachAI

## Project Goal

FaceitCoachAI is an AI-powered CS2 demo analysis platform.

Unlike traditional demo analyzers, the parser does NOT evaluate gameplay.

Its only responsibility is collecting structured facts.

The coaching and reasoning will be generated later by ChatGPT.

Architecture:

Demo
↓
Parser (Go)
↓
Facts
↓
AI Analysis
↓
React Dashboard

------------------------------------

## Technology Stack

Parser:
- Go

Frontend:
- React
- TypeScript

AI:
- ChatGPT API

------------------------------------

## Parser Principles

The parser NEVER makes coaching decisions.

The parser NEVER labels actions as:

- Good
- Bad
- Excellent
- Poor

Those are AI responsibilities.

The parser only stores facts.

------------------------------------

## Current Parser Modules

Completed:

✓ Player Parser

✓ Round Parser

✓ Kill Parser

✓ Damage Parser

✓ PlayerState Parser

✓ Support Analyzer V2

✓ Trade Analyzer

Working:

Rotation Analyzer

Planned:

Crossfire Analyzer

Utility Analyzer

Visibility Analyzer

Economy Analyzer

Communication Analyzer

------------------------------------

## PlayerState

Every player snapshot contains:

- Tick
- Round
- SteamID
- PlayerID

Position

- X
- Y
- Z

Health

- HP
- Armor
- Alive

View

- ViewYaw
- ViewPitch

Movement

- Velocity

------------------------------------

## Support Analyzer

For every death event:

Store:

Victim

Support Candidates

Each SupportCandidate contains:

Player

SteamID

PlayerID

X

Y

Z

Distance

HeightDifference

Alive

(Currently adding)

ViewYaw

ViewPitch

Later:

HP

Armor

Velocity

Candidates must be sorted by Distance.

Candidate[0] is the nearest teammate.

------------------------------------

## Trade Analyzer

For every trade:

Store:

Round

Tick

Victim

Killer

Trader

ExpectedTrader

TradeTimeTicks

SupportDistance

HeightDifference

The parser stores only facts.

AI decides whether a trade is good.

------------------------------------

## Rotation Analyzer

Planned.

Rotation should store:

Start Tick

End Tick

Start Position

End Position

Travel Distance

Travel Time

No judgement.

------------------------------------

## Crossfire Analyzer

Planned.

Store:

Angle

Distance

Position

Visibility

No judgement.

------------------------------------

## Visibility Analyzer

Future module.

Calculate:

Can teammate see enemy?

Can teammate see victim?

Angle to enemy

Line of sight

------------------------------------

## AI Layer

ChatGPT receives structured facts only.

Example:

Victim Position

Support Candidates

Trade Information

Visibility

Rotation

Then ChatGPT generates:

Mistakes

Suggestions

Coaching

Priority improvements

------------------------------------

## React Dashboard

Dashboard should show:

Timeline

Kills

Trades

Support

Rotation

Crossfire

Heatmap

MiniMap

Player Comparison

AI Coaching Report

------------------------------------

## Development Rules

1.

Never duplicate data.

2.

Always reuse PlayerState.

3.

Never parse demo twice.

4.

Parser only stores facts.

5.

No AI judgement inside parser.

6.

Every analyzer outputs structured data.

7.

Every analyzer should be reusable.

8.

Keep architecture modular.

9.

Support future AI reasoning.

10.

Performance is important.