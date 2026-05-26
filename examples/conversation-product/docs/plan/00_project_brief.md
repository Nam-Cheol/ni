# Project brief

## Product type

conversation_product

## Delivery surfaces

- conversation

## Purpose

Travel Concierge Triage is a human-operated conversation protocol. It gathers enough travel intent, constraints, and unresolved questions for a human concierge to continue planning safely.

## Problem

Travel intake often mixes hard constraints, preferences, and unsupported requests. The product needs a locked planning contract that proves ni can describe a non-software conversational product before any implementation or runtime exists.

## Success definition

The plan is successful when `ni status` reports READY, `ni end` locks the contract, and `ni run --target human-team` produces a handoff prompt grounded in conversation capabilities, risks, and transcript evaluations.
