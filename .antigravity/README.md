# AI Agent Context - Backend Interview Handbook

## Quick Start for AI Agents

If you're an AI agent resuming work on this project, start here:

1. **Read this file first** - You're doing it! ✅
2. **Check status**: Read `.antigravity/status.md` for current progress
4. **Learn patterns**: Read `.antigravity/implementation_guide.md` for how to add topics
5. **Understand architecture**: Read `.antigravity/architecture.md` for technical details
6. **Pick next task**: Check `status.md` for next priority topic
7. **Follow conventions**: Match existing code style and structure

## Project Summary (TL;DR)

**What**: Executable backend engineering interview preparation handbook (Go-powered)  
**Goal**: Help backend engineers prepare for interviews through runnable Go code examples covering CS fundamentals, Go internals, and system design  
**Status**: Fundamentals enriched with 25 topics (OS, Networking, Database), navigation system with 16 index files, 131 tests passing  
**Next**: Continue with concurrency and advanced topics

## Key Principles

1. **Executable Learning** - Every concept has runnable tests
2. **Code > Theory** - Show, don't tell
3. **Interview-Focused** - Content directly relevant to interviews
4. **Standard Library Only** - No external dependencies
5. **Isolation** - Each topic independent

## File Structure

```
.antigravity/
├── README.md                 # This file - Start here
├── project_overview.md       # High-level project info
├── implementation_guide.md   # How to add new topics
├── architecture.md           # Technical architecture
└── status.md                 # Current progress & next steps

Navigation System (NEW):
├── fundamentals/README.md    # Main fundamentals index
│   ├── os/README.md          # OS topics index
│   ├── networking/README.md  # Networking topics index
│   └── database/README.md    # Database topics index
└── internal/README.md        # Main internal packages index
    ├── basics/README.md      # Go basics index
    ├── intermediate/README.md
    ├── advanced/README.md
    ├── concurrency/README.md
    ├── memory/README.md
    ├── internals/README.md
    ├── patterns/README.md
    ├── system_design/README.md
    ├── ds/README.md
    ├── algo/README.md
    └── leetcode/README.md
```

## Current State

### ✅ Completed
- Basics package (6 topics, 82 tests)
- Partial concurrency (4 topics, 49 tests)
- Project infrastructure
- Documentation

### 🔄 In Progress
- Concurrency package (4/10 complete)

### ⏳ TODO
- 6 more concurrency topics
- 9 intermediate topics
- 8 advanced topics
- 7 memory topics
- 6 internals topics
- 7 patterns topics
- 8 data structures topics
- 11 algorithms topics
- 6 system design topics

**Total Remaining**: 60 topics

## Quick Reference

### Adding a New Topic

1. Create `internal/[package]/[topic].go`
2. Add required comment blocks (see implementation_guide.md)
3. Implement 5-10 example functions
4. Create `internal/[package]/[topic]_test.go`
5. Write comprehensive tests
6. Verify: `go test ./internal/[package]/ -v`
7. Check race: `go test -race ./internal/[package]/`
8. Update `status.md`

### Required Comment Blocks

Every `.go` file must have:
```go
// Why interviewers ask this:
// [explanation]

// Common pitfalls:
// [list of mistakes]

// Key takeaway:
// [main concept]
```

### Test Naming

Format: `Test[Topic]_[Behavior]`

Examples:
- `TestChannel_BufferedVsUnbuffered`
- `TestMutex_SafeConcurrentIncrement`
- `TestInterface_NilInterfaceGotcha`

### Quality Checklist

Before marking a topic complete:
- [ ] Required comment blocks present
- [ ] 5+ test functions
- [ ] Tests pass: `go test ./internal/[package]/`
- [ ] No races: `go test -race ./internal/[package]/`
- [ ] Vet passes: `go vet ./internal/[package]/`
- [ ] Coverage > 85%
- [ ] Deterministic tests (no `time.Sleep` in assertions)

## Common Commands

```bash
# Run all tests
go test ./...

# Run specific package
go test ./internal/concurrency/ -v

# Run with race detector
go test -race ./...

# Check coverage
go test -cover ./...

# Format code
go fmt ./...

# Vet code
go vet ./...

# Run test script
./scripts/run_all_tests.sh
```

## Priority Order

1. **Concurrency** (most important for interviews)
2. **Intermediate** (common interview topics)
3. **Advanced** (deeper knowledge)
4. **Patterns** (practical skills)
5. **Data Structures & Algorithms** (coding challenges)
6. **Memory & Internals** (performance understanding)
7. **System Design** (architectural thinking)

## Example Files to Reference

Good examples of proper implementation:
- `internal/basics/interfaces.go` - Interface patterns
- `internal/concurrency/channels.go` - Channel patterns
- `internal/concurrency/worker_pool.go` - Practical patterns
- `internal/basics/maps.go` - Gotcha documentation

## Known Issues

- `TestGoroutine_NoWaitCausesRace` intentionally demonstrates race condition
  - This is educational, not a bug
  - Fails with `-race` flag (expected)

## Important Notes

### Do's ✅
- Follow existing patterns
- Write deterministic tests
- Document gotchas
- Keep examples simple
- Focus on interview relevance
- Use descriptive names

### Don'ts ❌
- Add external dependencies
- Use `time.Sleep` in test assertions
- Create flaky tests
- Forget required comment blocks
- Make examples too complex
- Skip edge cases

## Next Immediate Task

**Complete Concurrency Package**

Add these 6 topics in order:
1. `channel_closing.go` - Channel closing rules
2. `select_stmt.go` - Advanced select patterns
3. `fan_in_out.go` - Fan-in/fan-out patterns
4. `atomic.go` - Atomic operations
5. `race_conditions.go` - Race detection/prevention
6. `context_goroutines.go` - Context-aware patterns

Each topic needs:
- Implementation file (~150-250 lines)
- Test file (~200-300 lines)
- 5-10 test functions
- Required comment blocks

## Success Metrics

- All tests passing
- Coverage > 85% per package
- No race conditions (except educational examples)
- Clear, executable documentation
- Interview-relevant content

## Questions?

If you're unsure about something:
1. Check existing implementations for patterns
2. Read `implementation_guide.md` for detailed guidance
3. Review `architecture.md` for design decisions

## Contact Information

This is a self-contained project designed for AI agent continuation. No human contact needed - all information is in the documentation.

---

**Remember**: The goal is to help engineers prepare for Go interviews through hands-on, executable examples. Keep it simple, clear, and focused on interview-relevant content.

**Good luck!** 🚀
