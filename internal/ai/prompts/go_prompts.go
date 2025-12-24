package prompts

import (
	"fmt"
	"strings"
)

func GenerateGoPrompt(layer, featureName, fields, routerType string) string {
	entity := featureName
	entityLower := strings.ToLower(featureName)

	switch layer {

	// =====================
	// MODEL
	// =====================
	case "Model":
		return fmt.Sprintf(`
You are a Go code generator.

Context:
- Clean Architecture backend
- Package: model
- This file defines ONLY data structure

Task:
Create a Go struct named '%s' in package 'model'.

Fields:
%s

Rules:
1. Use struct only (NO methods).
2. Use json tags in snake_case.
3. Use db tags matching column names.
4. Do NOT add validation or logic.
5. Output ONLY valid Go code.
`, entity, fields)

	// =====================
	// REPOSITORY
	// =====================
	case "Repository":
		return fmt.Sprintf(`
You are a Go code generator.

Context:
- Clean Architecture backend
- Package: repository
- Repository handles data persistence

Task:
Create repository skeleton for entity '%s'.

Rules:
1. Define interface '%sRepository'.
2. Methods: Create, GetByID, Update, Delete.
3. Each method accepts context.Context.
4. Use '*model.%s' as entity type.
5. Define struct '%sRepositoryImpl' with *sql.DB dependency.
6. Method bodies MUST be empty or return nil.
7. Do NOT write SQL queries.
8. Output ONLY valid Go code.
`, entity, entity, entity, entity)

	// =====================
	// SERVICE
	// =====================
	case "Service":
		return fmt.Sprintf(`
You are a Go code generator.

Context:
- Clean Architecture backend
- Package: service
- Service contains business abstraction

Task:
Create service skeleton for entity '%s'.

Rules:
1. Define interface '%sService'.
2. Methods: Create, GetByID, Update, Delete.
3. Each method accepts context.Context.
4. Use '*model.%s' as entity type.
5. Define struct '%sServiceImpl' with repository dependency.
6. Method bodies MUST be empty or return nil.
7. Do NOT implement business logic.
8. Output ONLY valid Go code.
`, entity, entity, entity, entity)

	// =====================
	// HANDLER
	// =====================
	case "Handler":
		return fmt.Sprintf(`
You are a Go code generator.

Context:
- Clean Architecture backend
- Package: handler
- Standard net/http handlers

Task:
Create HTTP handler skeleton for entity '%s'.

Rules:
1. Define struct '%sHandler' with service dependency.
2. Methods: Create, GetByID, Update, Delete.
3. Method signature must use (w http.ResponseWriter, r *http.Request).
4. Method bodies MUST be empty.
5. Do NOT write response logic.
6. Output ONLY valid Go code.
`, entity, entity)

	// =====================
	// ROUTER
	// =====================
	case "Router":
		var routerArg string
		var handlerType string
		var routeComment string

		switch strings.ToLower(routerType) {

		case "gin":
			routerArg = "*gin.Engine"
			handlerType = "*" + featureName + "Handler"
			routeComment = `
- POST   /%s
- GET    /%s/:id
- PUT    /%s/:id
- DELETE /%s/:id
`

		case "chi":
			routerArg = "chi.Router"
			handlerType = "*" + featureName + "Handler"
			routeComment = `
- POST   /%s
- GET    /%s/{id}
- PUT    /%s/{id}
- DELETE /%s/{id}
`

		default:
			routerArg = "*http.ServeMux"
			handlerType = "*" + featureName + "Handler"
			routeComment = `
- POST   /%s
- GET    /%s/{id}
- PUT    /%s/{id}
- DELETE /%s/{id}
`
		}

		return fmt.Sprintf(`
You are a Go code generator.

Context:
- Clean Architecture backend
- Package: router
- Router type: %s

Task:
Create router registration function for entity '%s'.

Rules:
1. Create function 'Register%sRoutes'.
2. Function accepts (%s, %s).
3. Inside the function, add ONLY comments for routes:
%s
4. Do NOT implement routing logic.
5. Do NOT import frameworks unless required by the signature.
6. Output ONLY valid Go code.
`,
			routerType,
			featureName,
			featureName,
			routerArg,
			handlerType,
			fmt.Sprintf(
				routeComment,
				entityLower, entityLower, entityLower, entityLower,
			),
		)

	}

	return ""
}
