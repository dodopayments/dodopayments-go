// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workloadidentity

import (
	"context"

	"github.com/dodopayments/dodopayments-go/internal/requestconfig"
	"github.com/dodopayments/dodopayments-go/option"
)

// GitHubActions requests a GitHub Actions OIDC token before each exchange. The
// job must have the "id-token: write" permission. An empty audience uses the
// origin of the client base URL, for example "https://live.dodopayments.com".
func GitHubActions(audience string) option.SubjectTokenProvider {
	return requestconfig.GitHubActionsSubjectTokenProvider(audience)
}

// GitLab reads a GitLab CI ID token from an environment variable. An empty name
// uses DODO_PAYMENTS_ID_TOKEN.
func GitLab(envVar string) option.SubjectTokenProvider {
	if envVar == "" {
		envVar = "DODO_PAYMENTS_ID_TOKEN"
	}
	return requestconfig.EnvSubjectTokenProvider(envVar)
}

// TerraformCloud reads TFC_WORKLOAD_IDENTITY_TOKEN_DODO or
// TFC_WORKLOAD_IDENTITY_TOKEN.
func TerraformCloud() option.SubjectTokenProvider {
	return requestconfig.TerraformCloudSubjectTokenProvider()
}

// GCP requests an identity token from the GCP metadata server before each
// exchange. An empty audience uses the origin of the client base URL.
func GCP(audience string) option.SubjectTokenProvider {
	return requestconfig.GCPSubjectTokenProvider(audience)
}

// Vercel exchanges the Vercel OIDC token for a token with the Dodo Payments
// audience before each exchange. It reads VERCEL_OIDC_TOKEN_FILE or
// VERCEL_OIDC_TOKEN. The audience is the origin of the client base URL.
func Vercel() option.SubjectTokenProvider {
	return requestconfig.VercelSubjectTokenProvider("")
}

// VercelWithAudience is the same as [Vercel], but it uses the audience that you
// supply. An empty audience uses the origin of the client base URL.
func VercelWithAudience(audience string) option.SubjectTokenProvider {
	return requestconfig.VercelSubjectTokenProvider(audience)
}

// Kubernetes reads a projected service account token before each exchange. An
// empty path uses /var/run/secrets/tokens/dodo-token.
func Kubernetes(path string) option.SubjectTokenProvider {
	if path == "" {
		path = requestconfig.DefaultKubernetesTokenPath
	}
	return requestconfig.FileSubjectTokenProvider(path)
}

// File reads the subject token from a file before each exchange. An empty path
// uses DODO_PAYMENTS_IDENTITY_TOKEN_FILE.
func File(path string) option.SubjectTokenProvider {
	return requestconfig.FileSubjectTokenProvider(path)
}

// AWS uses a function that returns a JWT, for example from STS
// GetWebIdentityToken. The function receives the default audience.
func AWS(fn func(ctx context.Context, audience string) (string, error)) option.SubjectTokenProvider {
	return option.SubjectTokenProviderFunc(fn)
}

// Static always returns the same token. Use it for tests.
func Static(token string) option.SubjectTokenProvider {
	return option.SubjectTokenProviderFunc(func(context.Context, string) (string, error) {
		return token, nil
	})
}

// Auto detects a source before each exchange. It examines GitHub Actions,
// Terraform Cloud, DODO_PAYMENTS_IDENTITY_TOKEN_FILE, and the Kubernetes token
// path, in that order.
func Auto() option.SubjectTokenProvider {
	return requestconfig.AutoSubjectTokenProvider()
}
