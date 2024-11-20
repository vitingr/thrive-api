# Thrive API

Thrive is a social networking platform designed to promote healthy habits such as running, diets, gym routines, and more. Users can track their health activities, join groups, follow friends, and share their progress with the community. This API powers the back-end of the Thrive platform, providing endpoints for managing users, activities, health habits, and social interactions.

## Table of Contents
- [Overview](#overview)
- [Authentication](#authentication)
- [API Endpoints](#api-endpoints)
  - [User Endpoints](#user-endpoints)
  - [Activity Endpoints](#activity-endpoints)
  - [Group Endpoints](#group-endpoints)
  - [Social Endpoints](#social-endpoints)
- [Data Models](#data-models)
- [Usage Example](#usage-example)
- [Error Handling](#error-handling)
- [Rate Limiting](#rate-limiting)
- [Contributing](#contributing)
- [License](#license)

## Overview
Thrive helps users track and share their health-related activities, connect with friends, join groups focused on health habits, and provide motivation to lead healthier lives. The API allows developers to interact with the platform's core features such as:
- User profiles and account management
- Activity tracking (e.g., runs, gym sessions, diet logs)
- Social networking features (e.g., follow users, join groups)
- Insights and analytics on health progress

## Authentication
All API requests require authentication via **OAuth 2.0** or **API tokens**.

### OAuth Authentication
- To use OAuth, obtain an authorization token by following the OAuth flow. 
- The token should be included in the `Authorization` header of the request as a Bearer token.

### API Token Authentication
- If you're using API tokens, include your API token in the request header like so:
