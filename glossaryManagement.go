package godeeplapi

import (
	"context"
	"github.com/AdolfZahid1/godeeplapi/models"
)

// ListLangPairsSupportedByGlossaries retrieves the list of language pairs supported
// by the glossary feature.
func (c *Client) ListLangPairsSupportedByGlossaries(ctx context.Context) (*models.GlossaryListResponse, error) {
	if err := c.checkAuth(); err != nil {
		return nil, err
	}

	respBody, err := c.doRequest(ctx, "GET", "/glossary-language-pairs", nil, nil)
	if err != nil {
		return nil, err
	}

	var resp models.GlossaryListResponse
	if err := unmarshalResponse(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// CreateGlossary creates a new glossary and returns info about it.
// Note: Use V3 client for this endpoint.
func (c *Client) CreateGlossary(ctx context.Context, req models.CreateGlossaryRequest) (*models.CreateGlossaryResponse, error) {
	if err := c.checkAuth(); err != nil {
		return nil, err
	}

	respBody, err := c.doRequest(ctx, "POST", "/glossaries", req, nil)
	if err != nil {
		return nil, err
	}

	var resp models.CreateGlossaryResponse
	if err := unmarshalResponse(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ListAllGlossaries returns all glossaries and their meta-information,
// but not the glossary entries.
// Note: Use V3 client for this endpoint.
func (c *Client) ListAllGlossaries(ctx context.Context) (*models.AllGlossaryListResponse, error) {
	if err := c.checkAuth(); err != nil {
		return nil, err
	}

	respBody, err := c.doRequest(ctx, "GET", "/glossaries", nil, nil)
	if err != nil {
		return nil, err
	}

	var resp models.AllGlossaryListResponse
	if err := unmarshalResponse(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// GetGlossaryByID retrieves meta information for a single glossary,
// omitting the glossary entries.
// Note: Use V3 client for this endpoint.
func (c *Client) GetGlossaryByID(ctx context.Context, id string) (*models.Glossary, error) {
	if err := c.checkAuth(); err != nil {
		return nil, err
	}

	respBody, err := c.doRequest(ctx, "GET", "/glossaries/"+id, nil, nil)
	if err != nil {
		return nil, err
	}

	var resp models.Glossary
	if err := unmarshalResponse(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// EditGlossary edits glossary details, such as name or a dictionary
// for a source and target language.
// Note: Use V3 client for this endpoint.
func (c *Client) EditGlossary(ctx context.Context, id string, req models.EditGlossaryRequest) (*models.Glossary, error) {
	if err := c.checkAuth(); err != nil {
		return nil, err
	}

	respBody, err := c.doRequest(ctx, "PATCH", "/glossaries/"+id, req, nil)
	if err != nil {
		return nil, err
	}

	var resp models.Glossary
	if err := unmarshalResponse(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// DeleteGlossary deletes the specified glossary.
// Note: Use V3 client for this endpoint.
func (c *Client) DeleteGlossary(ctx context.Context, id string) error {
	if err := c.checkAuth(); err != nil {
		return err
	}

	_, err := c.doRequest(ctx, "DELETE", "/glossaries/"+id, nil, nil)
	return err
}

// DeleteAllLangDictionaries deletes the dictionary associated with
// the given language pair with the given glossary ID.
// Note: Use V3 client for this endpoint.
func (c *Client) DeleteAllLangDictionaries(ctx context.Context, id string, query models.GlossaryLangPair) error {
	if err := c.checkAuth(); err != nil {
		return err
	}

	_, err := c.doRequestWithQuery(ctx, "DELETE", "/glossaries/"+id+"/dictionaries", nil, nil, query)
	return err
}

// GetGlossaryEntries lists the entries of a single glossary in tsv format.
// Note: Use V3 client for this endpoint.
func (c *Client) GetGlossaryEntries(ctx context.Context, id string, query models.GlossaryLangPair) (*models.GlossaryEntriesResponse, error) {
	if err := c.checkAuth(); err != nil {
		return nil, err
	}

	respBody, err := c.doRequestWithQuery(ctx, "GET", "/glossaries/"+id+"/entries", nil, nil, query)
	if err != nil {
		return nil, err
	}

	var resp models.GlossaryEntriesResponse
	if err := unmarshalResponse(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// ReplaceOrCreateDictionaryInGlossary replaces or creates a dictionary
// in the glossary with the specified entries.
// Note: Use V3 client for this endpoint.
func (c *Client) ReplaceOrCreateDictionaryInGlossary(ctx context.Context, id string, req models.Dictionary) (*models.EditOrCreateDictionaryInGlossaryResponse, error) {
	if err := c.checkAuth(); err != nil {
		return nil, err
	}

	respBody, err := c.doRequestWithQuery(ctx, "PUT", "/glossaries/"+id+"/dictionaries", req, nil, nil)
	if err != nil {
		return nil, err
	}

	var resp models.EditOrCreateDictionaryInGlossaryResponse
	if err := unmarshalResponse(respBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
