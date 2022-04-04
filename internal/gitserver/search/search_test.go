	"github.com/sourcegraph/sourcegraph/lib/errors"
		"echo consectetur adipiscing elit again > file3",
		require.Empty(t, matches[0].ModifiedFiles)

	t.Run("match both, in order with modified files", func(t *testing.T) {
		query := &protocol.MessageMatches{Expr: "c"}
		tree, err := ToMatchTree(query)
		require.NoError(t, err)
		searcher := &CommitSearcher{
			RepoDir:              dir,
			Query:                tree,
			IncludeModifiedFiles: true,
		}
		var matches []*protocol.CommitMatch
		err = searcher.Search(context.Background(), func(match *protocol.CommitMatch) {
			matches = append(matches, match)
		})
		require.NoError(t, err)
		require.Len(t, matches, 2)
		require.Equal(t, matches[0].Author.Name, "camden2")
		require.Equal(t, matches[1].Author.Name, "camden1")
		require.Equal(t, []string{"file2", "file3"}, matches[0].ModifiedFiles)
		require.Equal(t, []string{"file1"}, matches[1].ModifiedFiles)
	})
	}{
		{
			input: []byte(
				"\x1E2061ba96d63cba38f20a76f039cf29ef68736b8a\x00\x00HEAD\x00Camden Cheek\x00camden@sourcegraph.com\x001632251505\x00Camden Cheek\x00camden@sourcegraph.com\x001632251505\x00fix import\n\x005230097b75dcbb2c214618dd171da4053aff18a6\x00\x00" +
					"\x1E5230097b75dcbb2c214618dd171da4053aff18a6\x00\x00HEAD\x00Camden Cheek\x00camden@sourcegraph.com\x001632248499\x00Camden Cheek\x00camden@sourcegraph.com\x001632248499\x00only set matches if they exist\n\x00\x00",
			),
			expected: []*RawCommit{
				{
					Hash:           []byte("2061ba96d63cba38f20a76f039cf29ef68736b8a"),
					RefNames:       []byte(""),
					SourceRefs:     []byte("HEAD"),
					AuthorName:     []byte("Camden Cheek"),
					AuthorEmail:    []byte("camden@sourcegraph.com"),
					AuthorDate:     []byte("1632251505"),
					CommitterName:  []byte("Camden Cheek"),
					CommitterEmail: []byte("camden@sourcegraph.com"),
					CommitterDate:  []byte("1632251505"),
					Message:        []byte("fix import"),
					ParentHashes:   []byte("5230097b75dcbb2c214618dd171da4053aff18a6"),
					ModifiedFiles:  [][]byte{{}, {}},
				},
				{
					Hash:           []byte("5230097b75dcbb2c214618dd171da4053aff18a6"),
					RefNames:       []byte(""),
					SourceRefs:     []byte("HEAD"),
					AuthorName:     []byte("Camden Cheek"),
					AuthorEmail:    []byte("camden@sourcegraph.com"),
					AuthorDate:     []byte("1632248499"),
					CommitterName:  []byte("Camden Cheek"),
					CommitterEmail: []byte("camden@sourcegraph.com"),
					CommitterDate:  []byte("1632248499"),
					Message:        []byte("only set matches if they exist"),
					ParentHashes:   []byte(""),
					ModifiedFiles:  [][]byte{{}},
				},
			},
		},
		{
			input: []byte(
				"\x1E2061ba96d63cba38f20a76f039cf29ef68736b8a\x00\x00HEAD\x00Camden Cheek\x00camden@sourcegraph.com\x001632251505\x00Camden Cheek\x00camden@sourcegraph.com\x001632251505\x00fix import\n\x005230097b75dcbb2c214618dd171da4053aff18a6\x00\x00file1" +
					"\x1E5230097b75dcbb2c214618dd171da4053aff18a6\x00\x00HEAD\x00Camden Cheek\x00camden@sourcegraph.com\x001632248499\x00Camden Cheek\x00camden@sourcegraph.com\x001632248499\x00only set matches if they exist\n\x00\x00file1\x00file2",
			),
			expected: []*RawCommit{
				{
					Hash:           []byte("2061ba96d63cba38f20a76f039cf29ef68736b8a"),
					RefNames:       []byte(""),
					SourceRefs:     []byte("HEAD"),
					AuthorName:     []byte("Camden Cheek"),
					AuthorEmail:    []byte("camden@sourcegraph.com"),
					AuthorDate:     []byte("1632251505"),
					CommitterName:  []byte("Camden Cheek"),
					CommitterEmail: []byte("camden@sourcegraph.com"),
					CommitterDate:  []byte("1632251505"),
					Message:        []byte("fix import"),
					ParentHashes:   []byte("5230097b75dcbb2c214618dd171da4053aff18a6"),
					ModifiedFiles: [][]byte{
						{},
						[]byte("file1"),
					},
				},
				{
					Hash:           []byte("5230097b75dcbb2c214618dd171da4053aff18a6"),
					RefNames:       []byte(""),
					SourceRefs:     []byte("HEAD"),
					AuthorName:     []byte("Camden Cheek"),
					AuthorEmail:    []byte("camden@sourcegraph.com"),
					AuthorDate:     []byte("1632248499"),
					CommitterName:  []byte("Camden Cheek"),
					CommitterEmail: []byte("camden@sourcegraph.com"),
					CommitterDate:  []byte("1632248499"),
					Message:        []byte("only set matches if they exist"),
					ParentHashes:   []byte(""),
					ModifiedFiles: [][]byte{
						[]byte("file1"),
						[]byte("file2"),
					},
				},
			},
		},
	}