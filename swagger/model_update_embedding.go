/*
 * FastAPI
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type UpdateEmbedding struct {
	Embeddings     []Object      `json:"embeddings,omitempty"`
	Metadatas      []interface{} `json:"metadatas,omitempty"`
	Documents      []string      `json:"documents,omitempty"`
	Ids            []string      `json:"ids"`
	IncrementIndex bool          `json:"increment_index,omitempty"`
}
