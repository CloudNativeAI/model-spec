/*
 *     Copyright 2024 The CNAI Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package v1

const (
	// ArtifactTypeModelManifest specifies the media type for a model manifest.
	ArtifactTypeModelManifest = "application/vnd.cnai.model.manifest.v1+json"
)

const (
	// MediaTypeModelLayer is the media type used for layers referenced by the manifest, it's recommended to use
	// the `application/vnd.oci.image.layer.v1.tar` format for the model layer to avoid compression.
	MediaTypeModelLayer = "application/vnd.cnai.model.layer.v1.tar"

	// MediaTypeModelLayerGzip is the media type used for gzipped layers
	// referenced by the manifest.
	MediaTypeModelLayerGzip = "application/vnd.cnai.model.layer.v1.tar+gzip"

	// MediaTypeModelDoc is the media type used for documentation referenced by the manifest,
	// such as README.md, LICENSE, etc.
	MediaTypeModelDoc = "application/vnd.cnai.model.doc.v1.tar"

	// MediaTypeModelConfig is the media type used for configuration referenced by the manifest,
	// such as config.json, tokenizer.json, generation_config.json, etc.
	MediaTypeModelConfig = "application/vnd.cnai.model.config.v1.tar"
)
