# Contributing

Thank you for your interest in contributing to Hammer!

## Workflow

1. **Open an issue first**
   - Create an issue describing the bug, improvement, or new feature you want to work on.
   - Clearly state the motivation and, if applicable, steps to reproduce (for bugs).

2. **Fork and clone the repository**
   On GitHub, fork `sergioc0sta/hammer` to your account, then:

   ```bash
   git clone https://github.com/sergioc0sta/hammer.git
   cd hammer
   ```

3. **Create a feature branch from main**
   Use the following naming convention, where X is the related issue number:

   ```bash
   git checkout -b feat/hmm-X/short-description
   # example:
   # git checkout -b feat/hmm-12/add-json-report
   ```

4. **Implement your changes**
   - Keep commits small and focused.
   - Write clear commit messages.
   - Update or add tests when relevant.

5. **Run tests before submitting**

   ```bash
   go test ./...
   ```

6. **Push your branch and open a Pull Request**

   ```bash
   git push origin feat/hmm-X/short-description
   ```

   Then, on GitHub:

   - Open a Pull Request from your branch to `sergioc0sta/hammer`’s `main` branch.
   - Reference the related issue in the PR description (for example: Fixes #X).
   - Briefly explain:
     - What you changed
     - Why you changed it
     - How to test it

We appreciate every contribution, from small fixes to new features.
