# Trooper

**Trooper** is an opinionated, YAML-first infrastructure and DevOps platform that standardizes how cloud resources are defined, reviewed, and deployed at scale.

Trooper’s goal is simple but ambitious:

> **Eliminate the operational chaos caused by ad-hoc Terraform, inconsistent pipelines, and tribal DevOps knowledge — without slowing teams down.**

Trooper achieves this by letting teams declare *what they want* in a single YAML file, while the platform handles *how it is built, validated, and applied* using Terraform and proven best practices.

---

## What Trooper Is (and Is Not)

### Trooper **is**
- A **platform abstraction** over Terraform (and optionally Terragrunt)
- A **YAML-driven resource catalog** model
- A **standardization layer** for cloud infra, CI/CD, guardrails, and governance
- A way for teams to self-serve infrastructure safely

### Trooper **is not**
- A replacement for Terraform
- A new cloud provider
- A proprietary runtime that locks you in
- A “magic” black box — everything is auditable and deterministic

---

## Core Philosophy

### 1. YAML is the Single Source of Truth
Teams define infrastructure intent in **one file**: `troop.yaml`. 

Everything else is generated and executed by Trooper.

### 2. Platform Teams Own the “How”
Application teams focus on declaring resources, while the platform team manages:
- Terraform backends and Provider pinning
- IAM complexity and CI/CD wiring
- Policy enforcement

### 3. Deterministic, Auditable Execution
- Generates Terraform at runtime in an isolated workdir
- Pins module and provider versions
- Runs policy checks before plan/apply
- Same YAML + same versions = same result

### 4. Clean Repos by Default
A Trooper-managed repo is intentionally minimal:
- troop.yaml
- .github/workflows/
- README.md

---

## The troop CLI
`troop` is the command-line interface for interacting with the platform.

- **troop init**: Initialize a project and CI workflows
- **troop validate**: Validate schema and intent
- **troop plan --env dev**: Generate and review changes
- **troop apply --env prod**: Execute changes safely

---

## Typical Workflow

1. **Initialize a Project**: Run `troop init`.
2. **Define Resources**: Add resources to `troop.yaml`.
3. **Review via Pull Request**: PR runs `troop plan` and policy checks.
4. **Apply via CI**: Merge triggers `troop apply`.

---

## Architecture Overview

Trooper is intentionally layered:
**troop.yaml (intent) -> Trooper CLI (generation) -> Terraform (execution) -> Cloud Provider**

### Module System
Each YAML resource type maps to a versioned Terraform module:
- **network.vnet** -> azure/network/vnet
- **k8s.aks** -> azure/kubernetes/aks
- **security.keyVault** -> azure/security/keyvault

---

## State & Governance
- **State Management**: Scaffolds remote state (encrypted/locked) by default. State is never committed to the repo.
- **Guardrails**: Designed to support Policy-as-code (OPA, Checkov), naming standards, and cost controls.

---

## Mission
Make infrastructure boring, safe, and repeatable — without slowing innovation.

---

## Status
Trooper is under active development. Interfaces and schemas may evolve, but the core principles are stable.