import { TestBed, inject } from '@angular/core/testing';

import { ArticlePostService } from './article-post.service';

describe('ArticlePostService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ArticlePostService]
    });
  });

  it('should ...', inject([ArticlePostService], (service: ArticlePostService) => {
    expect(service).toBeTruthy();
  }));
});
