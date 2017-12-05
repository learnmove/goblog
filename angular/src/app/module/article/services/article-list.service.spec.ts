import { TestBed, inject } from '@angular/core/testing';

import { ArticleListService } from './article-list.service';

describe('ArticleListService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [ArticleListService]
    });
  });

  it('should ...', inject([ArticleListService], (service: ArticleListService) => {
    expect(service).toBeTruthy();
  }));
});
